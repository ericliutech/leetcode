// Package judge runs shared JSON fixtures against typed solution functions.
package judge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

type Case struct {
	ID       int                        `json:"id"`
	Input    map[string]json.RawMessage `json:"input"`
	Expected json.RawMessage            `json:"expected"`
	Remark   string                     `json:"remark"`
}

// Load rejects incomplete fixtures instead of treating absent values as zero.
func Load(path string) ([]Case, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cases []Case
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cases); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		return nil, fmt.Errorf("%s: unexpected content after case array", path)
	}
	if len(cases) == 0 {
		return nil, fmt.Errorf("%s: no test cases", path)
	}
	ids := make(map[int]bool)
	for _, tc := range cases {
		if tc.ID <= 0 || ids[tc.ID] {
			return nil, fmt.Errorf("%s: IDs must be positive and unique: %d", path, tc.ID)
		}
		ids[tc.ID] = true
		if tc.Input == nil || len(tc.Expected) == 0 {
			return nil, fmt.Errorf("%s case %d: input and expected are required", path, tc.ID)
		}
	}
	return cases, nil
}

func decode(data json.RawMessage, kind reflect.Type) (reflect.Value, error) {
	if bytes.Equal(bytes.TrimSpace(data), []byte("null")) {
		switch kind.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Interface:
		default:
			return reflect.Value{}, fmt.Errorf("null is not a valid %s", kind)
		}
	}
	if kind.Kind() == reflect.Array {
		var elements []json.RawMessage
		if err := json.Unmarshal(data, &elements); err != nil || len(elements) != kind.Len() {
			return reflect.Value{}, fmt.Errorf("expected exactly %d elements for %s", kind.Len(), kind)
		}
	}
	value := reflect.New(kind)
	if err := json.Unmarshal(data, value.Interface()); err != nil {
		return reflect.Value{}, fmt.Errorf("cannot decode %s as %s: %w", data, kind, err)
	}
	return value.Elem(), nil
}

func arguments(tc Case, signature reflect.Type, names []string) ([]reflect.Value, error) {
	if len(tc.Input) != len(names) {
		return nil, fmt.Errorf("input keys must match parameters %v", names)
	}
	args := make([]reflect.Value, len(names))
	for index, name := range names {
		data, exists := tc.Input[name]
		if !exists {
			return nil, fmt.Errorf("missing input parameter %q", name)
		}
		value, err := decode(data, signature.In(index))
		if err != nil {
			return nil, fmt.Errorf("parameter %s: %w", name, err)
		}
		args[index] = value
	}
	return args, nil
}

func matches(got, expected reflect.Value, comparison string) (bool, error) {
	switch comparison {
	case "exact":
		return reflect.DeepEqual(got.Interface(), expected.Interface()), nil
	case "unordered":
		if got.Kind() != reflect.Array && got.Kind() != reflect.Slice {
			return false, fmt.Errorf("unordered comparison requires an array or slice")
		}
		if got.Kind() == reflect.Slice && got.IsNil() != expected.IsNil() {
			return false, nil
		}
		if got.Len() != expected.Len() {
			return false, nil
		}
		// Compare only the outer collection; preserve multiplicity and inner order.
		used := make([]bool, expected.Len())
		for index := 0; index < got.Len(); index++ {
			found := false
			for other := 0; other < expected.Len(); other++ {
				if !used[other] && reflect.DeepEqual(got.Index(index).Interface(), expected.Index(other).Interface()) {
					used[other], found = true, true
					break
				}
			}
			if !found {
				return false, nil
			}
		}
		return true, nil
	default:
		return false, fmt.Errorf("unknown comparison %q", comparison)
	}
}

// Run is called by generated adapters. JSON keys are matched to parameter names,
// while reflection supplies the declared types and invokes the function.
func Run(t *testing.T, path string, function any, names []string, comparison string) {
	t.Helper()
	fn := reflect.ValueOf(function)
	if !fn.IsValid() || fn.Kind() != reflect.Func || fn.IsNil() {
		t.Fatal("adapter must supply a solution function")
	}
	signature := fn.Type()
	if signature.IsVariadic() || signature.NumIn() != len(names) || signature.NumOut() != 1 {
		t.Fatal("supported functions have named fixed parameters and exactly one result")
	}
	if comparison != "exact" && comparison != "unordered" {
		t.Fatalf("unknown comparison %q", comparison)
	}
	cases, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("case_%d", tc.ID), func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Errorf("panic: %v; input=%v expected=%s remark=%q", recovered, tc.Input, tc.Expected, tc.Remark)
				}
			}()
			args, err := arguments(tc, signature, names)
			if err != nil {
				t.Fatal(err)
			}
			want, err := decode(tc.Expected, signature.Out(0))
			if err != nil {
				t.Fatalf("expected: %v", err)
			}
			// Decode afresh for each call, so mutated inputs cannot leak between cases.
			got := fn.Call(args)[0]
			ok, err := matches(got, want, comparison)
			if err != nil {
				t.Fatal(err)
			}
			if !ok {
				t.Errorf("input=%v: got %v, want %v (%s); remark=%q",
					tc.Input, got.Interface(), want.Interface(), comparison, tc.Remark)
			}
		})
	}
}
