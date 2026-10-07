package judge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func fixture(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "testcases.json")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLoadRejectsInvalidFixtures(t *testing.T) {
	for name, data := range map[string]string{
		"empty":            `[]`,
		"malformed":        `[`,
		"duplicate IDs":    `[{"id":1,"input":{},"expected":0},{"id":1,"input":{},"expected":0}]`,
		"missing input":    `[{"id":1,"expected":0}]`,
		"missing expected": `[{"id":1,"input":{}}]`,
		"missing ID":       `[{"input":{},"expected":0}]`,
		"unknown field":    `[{"id":1,"input":{},"expected":0,"expecetd":1}]`,
		"trailing JSON":    `[{"id":1,"input":{},"expected":0}] {}`,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Load(fixture(t, data)); err == nil {
				t.Fatal("invalid fixture was accepted")
			}
		})
	}
}

func TestZeroAndFalseExpectedArePresent(t *testing.T) {
	cases, err := Load(fixture(t, `[{"id":1,"input":{"s":""},"expected":0},{"id":2,"input":{"s":"("},"expected":false}]`))
	if err != nil || len(cases) != 2 {
		t.Fatalf("zero and false must remain valid expectations: %v", err)
	}
	if cases[0].Remark != "" {
		t.Fatal("remark should default to blank")
	}
}

func TestArgumentsFollowSignatureOrder(t *testing.T) {
	function := func(first int, second string) string { return second }
	args, err := arguments(Case{Input: map[string]json.RawMessage{
		"second": json.RawMessage(`"hello"`), "first": json.RawMessage(`42`),
	}}, reflect.TypeOf(function), []string{"first", "second"})
	if err != nil {
		t.Fatal(err)
	}
	if args[0].Int() != 42 || args[1].String() != "hello" {
		t.Fatal("arguments did not follow signature order")
	}
	for _, input := range []map[string]json.RawMessage{
		{"first": json.RawMessage(`42`)},
		{"first": json.RawMessage(`42`), "wrong": json.RawMessage(`"hello"`)},
		{"first": json.RawMessage(`"not an integer"`), "second": json.RawMessage(`"hello"`)},
	} {
		if _, err := arguments(Case{Input: input}, reflect.TypeOf(function), []string{"first", "second"}); err == nil {
			t.Fatalf("invalid parameters accepted: %v", input)
		}
	}
}

func TestDecodePreservesTypes(t *testing.T) {
	value, err := decode(json.RawMessage(`9007199254740993`), reflect.TypeOf(int64(0)))
	if err != nil || value.Int() != 9007199254740993 {
		t.Fatalf("integer precision lost: %v", err)
	}
	if _, err := decode(json.RawMessage(`null`), reflect.TypeOf(0)); err == nil {
		t.Fatal("null integer accepted as zero")
	}
	if _, err := decode(json.RawMessage(`[1,2,3]`), reflect.TypeOf([2]int{})); err == nil {
		t.Fatal("oversized fixed array silently truncated")
	}
}

func TestComparisonsPreserveMultiplicityAndNestedOrder(t *testing.T) {
	checks := []struct {
		name, mode    string
		got, expected any
		want          bool
	}{
		{"either index order", "unordered", []int{1, 0}, []int{0, 1}, true},
		{"exact order", "exact", []int{1, 0}, []int{0, 1}, false},
		{"duplicate indices", "unordered", []int{0, 0}, []int{0, 1}, false},
		{"duplicate counts", "unordered", []int{1, 2, 2}, []int{1, 1, 2}, false},
		{"nested order", "unordered", [][]int{{2, 1}}, [][]int{{1, 2}}, false},
		{"null vs empty", "unordered", []int(nil), []int{}, false},
	}
	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			got, err := matches(reflect.ValueOf(check.got), reflect.ValueOf(check.expected), check.mode)
			if err != nil || got != check.want {
				t.Fatalf("match=%v, want %v: %v", got, check.want, err)
			}
		})
	}
	if _, err := matches(reflect.ValueOf(1), reflect.ValueOf(1), "unordered"); err == nil {
		t.Fatal("unordered scalar accepted")
	}
}

func TestRunDecodesFreshMutableInputs(t *testing.T) {
	path := fixture(t, `[{"id":1,"input":{"values":[7]},"expected":7},{"id":2,"input":{"values":[7]},"expected":7}]`)
	Run(t, path, func(values []int) int {
		original := values[0]
		values[0] = 99
		return original
	}, []string{"values"}, "exact")
}
