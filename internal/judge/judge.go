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
	trimmed := bytes.TrimSpace(data)
	if isListNode(kind) && (bytes.HasPrefix(trimmed, []byte("[")) || bytes.Equal(trimmed, []byte("null"))) {
		return decodeList(data, kind)
	}
	if isTreeNode(kind) && (bytes.HasPrefix(trimmed, []byte("[")) || bytes.Equal(trimmed, []byte("null"))) {
		return decodeTree(data, kind)
	}
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

// Recognize ordinary singly linked nodes without depending on their package.
// Extra fields (such as Random) require a different fixture representation.
func isListNode(kind reflect.Type) bool {
	if kind.Kind() != reflect.Pointer || kind.Elem().Kind() != reflect.Struct || kind.Elem().NumField() != 2 {
		return false
	}
	value, hasValue := kind.Elem().FieldByName("Val")
	next, hasNext := kind.Elem().FieldByName("Next")
	return hasValue && hasNext && len(value.Index) == 1 && len(next.Index) == 1 &&
		value.PkgPath == "" && next.PkgPath == "" && value.Type.Kind() == reflect.Int && next.Type == kind
}

func decodeList(data json.RawMessage, kind reflect.Type) (reflect.Value, error) {
	var elements []json.RawMessage
	if err := json.Unmarshal(data, &elements); err != nil {
		return reflect.Value{}, fmt.Errorf("linked list requires an array of integers: %w", err)
	}
	head := reflect.Zero(kind)
	for index := len(elements) - 1; index >= 0; index-- {
		node := reflect.New(kind.Elem()).Convert(kind)
		value, err := decode(elements[index], node.Elem().FieldByName("Val").Type())
		if err != nil {
			return reflect.Value{}, fmt.Errorf("linked-list element %d: %w", index, err)
		}
		node.Elem().FieldByName("Val").Set(value)
		node.Elem().FieldByName("Next").Set(head)
		head = node
	}
	return head, nil
}

func listValues(head reflect.Value) ([]int, error) {
	values := make([]int, 0)
	visited := make(map[any]bool)
	for !head.IsNil() {
		identity := head.Interface()
		if visited[identity] {
			return nil, fmt.Errorf("returned linked list contains a cycle")
		}
		visited[identity] = true
		values = append(values, int(head.Elem().FieldByName("Val").Int()))
		head = head.Elem().FieldByName("Next")
	}
	return values, nil
}

func display(value reflect.Value) any {
	if isListNode(value.Type()) {
		values, _ := listValues(value) // Comparison has already checked for cycles.
		return values
	}
	if isTreeNode(value.Type()) {
		values, _ := treeValues(value) // Comparison has already checked the shape.
		return values
	}
	return value.Interface()
}

func isTreeNode(kind reflect.Type) bool {
	if kind.Kind() != reflect.Pointer || kind.Elem().Kind() != reflect.Struct || kind.Elem().NumField() != 3 {
		return false
	}
	value, hasValue := kind.Elem().FieldByName("Val")
	left, hasLeft := kind.Elem().FieldByName("Left")
	right, hasRight := kind.Elem().FieldByName("Right")
	return hasValue && hasLeft && hasRight && value.Type.Kind() == reflect.Int &&
		len(value.Index) == 1 && len(left.Index) == 1 && len(right.Index) == 1 &&
		value.PkgPath == "" && left.PkgPath == "" && right.PkgPath == "" && left.Type == kind && right.Type == kind
}

func decodeTree(data json.RawMessage, kind reflect.Type) (reflect.Value, error) {
	var elements []json.RawMessage
	if err := json.Unmarshal(data, &elements); err != nil {
		return reflect.Value{}, fmt.Errorf("binary tree requires a level-order array: %w", err)
	}
	nodeAt := func(index int) (reflect.Value, error) {
		if bytes.Equal(bytes.TrimSpace(elements[index]), []byte("null")) {
			return reflect.Zero(kind), nil
		}
		node := reflect.New(kind.Elem()).Convert(kind)
		value, err := decode(elements[index], node.Elem().FieldByName("Val").Type())
		if err != nil {
			return reflect.Value{}, fmt.Errorf("tree element %d: %w", index, err)
		}
		node.Elem().FieldByName("Val").Set(value)
		return node, nil
	}
	if len(elements) == 0 {
		return reflect.Zero(kind), nil
	}
	root, err := nodeAt(0)
	if err != nil {
		return reflect.Value{}, err
	}
	var parents []reflect.Value
	if !root.IsNil() {
		parents = append(parents, root)
	}
	next := 1
	for index := 0; index < len(parents) && next < len(elements); index++ {
		for _, field := range []string{"Left", "Right"} {
			if next == len(elements) {
				break
			}
			child, err := nodeAt(next)
			if err != nil {
				return reflect.Value{}, err
			}
			next++
			parents[index].Elem().FieldByName(field).Set(child)
			if !child.IsNil() {
				parents = append(parents, child)
			}
		}
	}
	for ; next < len(elements); next++ {
		if !bytes.Equal(bytes.TrimSpace(elements[next]), []byte("null")) {
			return reflect.Value{}, fmt.Errorf("tree element %d has no parent", next)
		}
	}
	return root, nil
}

func treeValues(root reflect.Value) ([]any, error) {
	values := make([]any, 0)
	queue := []reflect.Value{root}
	visited := make(map[any]bool)
	for index := 0; index < len(queue); index++ {
		node := queue[index]
		if node.IsNil() {
			values = append(values, nil)
			continue
		}
		identity := node.Interface()
		if visited[identity] {
			return nil, fmt.Errorf("returned binary tree contains a cycle or shared node")
		}
		visited[identity] = true
		values = append(values, int(node.Elem().FieldByName("Val").Int()))
		queue = append(queue, node.Elem().FieldByName("Left"), node.Elem().FieldByName("Right"))
	}
	for len(values) > 0 && values[len(values)-1] == nil {
		values = values[:len(values)-1]
	}
	return values, nil
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
	if isTreeNode(got.Type()) {
		if comparison != "exact" {
			return false, fmt.Errorf("binary trees require exact structural comparison")
		}
		values, err := treeValues(got)
		if err != nil {
			return false, err
		}
		wanted, err := treeValues(expected)
		if err != nil {
			return false, err
		}
		return reflect.DeepEqual(values, wanted), nil
	}
	if isListNode(got.Type()) {
		if comparison != "exact" {
			return false, fmt.Errorf("linked lists require exact ordered comparison")
		}
		values, err := listValues(got)
		if err != nil {
			return false, err
		}
		wanted, err := listValues(expected)
		if err != nil {
			return false, err
		}
		return reflect.DeepEqual(values, wanted), nil
	}
	switch comparison {
	case "exact":
		return reflect.DeepEqual(got.Interface(), expected.Interface()), nil
	case "unordered", "combinations":
		if got.Kind() != reflect.Array && got.Kind() != reflect.Slice {
			return false, fmt.Errorf("unordered comparison requires an array or slice")
		}
		if comparison == "combinations" {
			inner := got.Type().Elem().Kind()
			if inner != reflect.Array && inner != reflect.Slice {
				return false, fmt.Errorf("combinations comparison requires nested arrays or slices")
			}
		}
		if got.Kind() == reflect.Slice && got.IsNil() != expected.IsNil() {
			return false, nil
		}
		if got.Len() != expected.Len() {
			return false, nil
		}
		// Match each entry once, preserving multiplicity at both levels.
		used := make([]bool, expected.Len())
		for index := 0; index < got.Len(); index++ {
			found := false
			for other := 0; other < expected.Len(); other++ {
				if used[other] {
					continue
				}
				equal := reflect.DeepEqual(got.Index(index).Interface(), expected.Index(other).Interface())
				if comparison == "combinations" {
					var err error
					equal, err = matches(got.Index(index), expected.Index(other), "unordered")
					if err != nil {
						return false, err
					}
				}
				if equal {
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
	if comparison != "exact" && comparison != "unordered" && comparison != "combinations" {
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
				t.Fatalf("input=%v: %v; remark=%q", tc.Input, err, tc.Remark)
			}
			if !ok {
				t.Errorf("input=%v: got %v, want %v (%s); remark=%q",
					tc.Input, display(got), display(want), comparison, tc.Remark)
			}
		})
	}
}
