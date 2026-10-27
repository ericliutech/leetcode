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
		{"combination order", "combinations", [][]int{{7}, {3, 2, 2}}, [][]int{{2, 2, 3}, {7}}, true},
		{"combination duplicates", "combinations", [][]int{{2, 2, 3}, {3, 2, 2}}, [][]int{{2, 2, 3}, {7}}, false},
		{"combination value counts", "combinations", [][]int{{2, 3, 3}}, [][]int{{2, 2, 3}}, false},
		{"missing combination", "combinations", [][]int{{7}}, [][]int{{7}, {2, 2, 3}}, false},
		{"empty combinations", "combinations", [][]int{}, [][]int{}, true},
		{"nil combinations", "combinations", [][]int(nil), [][]int(nil), true},
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

type testNode struct {
	Val  int
	Next *testNode
}

func TestLinkedListDecoding(t *testing.T) {
	kind := reflect.TypeOf((*testNode)(nil))
	for _, raw := range []string{`[]`, `null`, `[7]`, `[1,2,2,-3]`} {
		t.Run(raw, func(t *testing.T) {
			head, err := decode(json.RawMessage(raw), kind)
			if err != nil {
				t.Fatal(err)
			}
			got, err := listValues(head)
			if err != nil {
				t.Fatal(err)
			}
			want := make([]int, 0)
			if err := json.Unmarshal([]byte(raw), &want); err != nil {
				t.Fatal(err)
			}
			if raw == `null` {
				want = []int{}
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
			if len(want) == 0 && !head.IsNil() {
				t.Fatal("empty list must produce a nil head")
			}
		})
	}
	for _, raw := range []string{`[null]`, `[1.5]`, `["1"]`, `[[1]]`} {
		if _, err := decode(json.RawMessage(raw), kind); err == nil {
			t.Fatalf("invalid node values accepted: %s", raw)
		}
	}
}

func TestLinkedListDecodingCreatesIndependentNodes(t *testing.T) {
	kind := reflect.TypeOf((*testNode)(nil))
	first, err := decode(json.RawMessage(`[1,2]`), kind)
	if err != nil {
		t.Fatal(err)
	}
	second, err := decode(json.RawMessage(`[1,2]`), kind)
	if err != nil {
		t.Fatal(err)
	}
	first.Interface().(*testNode).Next.Val = 99
	values, err := listValues(second)
	if err != nil || !reflect.DeepEqual(values, []int{1, 2}) {
		t.Fatalf("decoded lists share mutable nodes: %v (%v)", values, err)
	}
}

func TestLinkedListComparisonPreservesOrderAndLength(t *testing.T) {
	kind := reflect.TypeOf((*testNode)(nil))
	for _, check := range []struct {
		got, want string
		equal     bool
	}{
		{`[3,2,1]`, `[3,2,1]`, true},
		{`[1,2,3]`, `[3,2,1]`, false},
		{`[1,2]`, `[1,2,3]`, false},
		{`[]`, `null`, true},
		{`[1,1]`, `[1,2]`, false},
	} {
		got, err := decode(json.RawMessage(check.got), kind)
		if err != nil {
			t.Fatal(err)
		}
		want, err := decode(json.RawMessage(check.want), kind)
		if err != nil {
			t.Fatal(err)
		}
		matched, err := matches(got, want, "exact")
		if err != nil || matched != check.equal {
			t.Fatalf("%s vs %s: matched=%v, want %v (%v)", check.got, check.want, matched, check.equal, err)
		}
	}
	if _, err := matches(reflect.ValueOf((*testNode)(nil)), reflect.ValueOf((*testNode)(nil)), "unordered"); err == nil {
		t.Fatal("linked lists must preserve order")
	}
}

func TestLinkedListCyclesAreReported(t *testing.T) {
	for _, twoNodes := range []bool{false, true} {
		head := &testNode{Val: 1}
		if twoNodes {
			head.Next = &testNode{Val: 2, Next: head}
		} else {
			head.Next = head
		}
		if _, err := matches(reflect.ValueOf(head), reflect.ValueOf((*testNode)(nil)), "exact"); err == nil {
			t.Fatal("cycle was not reported")
		}
	}
}

func TestLinkedListRecognitionDoesNotIgnoreAdditionalFields(t *testing.T) {
	type randomNode struct {
		Val    int
		Next   *randomNode
		Random *randomNode
	}
	if isListNode(reflect.TypeOf((*randomNode)(nil))) {
		t.Fatal("nodes with extra fields require their own representation")
	}
	if isListNode(reflect.TypeOf(testNode{})) || isListNode(reflect.TypeOf([]int{})) {
		t.Fatal("only pointers to ordinary singly linked nodes should match")
	}
}

type testTreeNode struct {
	Val   int
	Left  *testTreeNode
	Right *testTreeNode
}

func TestBinaryTreeLevelOrderDecoding(t *testing.T) {
	kind := reflect.TypeOf((*testTreeNode)(nil))
	for _, check := range []struct{ input, canonical string }{
		{`[]`, `[]`},
		{`null`, `[]`},
		{`[null]`, `[]`},
		{`[1]`, `[1]`},
		{`[1,null,null,null]`, `[1]`},
		{`[3,9,20,null,null,15,7]`, `[3,9,20,null,null,15,7]`},
		{`[1,null,2,3]`, `[1,null,2,3]`},
		{`[0,-1000,1000]`, `[0,-1000,1000]`},
	} {
		t.Run(check.input, func(t *testing.T) {
			root, err := decode(json.RawMessage(check.input), kind)
			if err != nil {
				t.Fatal(err)
			}
			values, err := treeValues(root)
			if err != nil {
				t.Fatal(err)
			}
			canonical, err := json.Marshal(values)
			if err != nil || string(canonical) != check.canonical {
				t.Fatalf("got %s, want %s (%v)", canonical, check.canonical, err)
			}
		})
	}
	root, err := decode(json.RawMessage(`[1,null,2,3]`), kind)
	if err != nil {
		t.Fatal(err)
	}
	node := root.Interface().(*testTreeNode)
	if node.Left != nil || node.Right.Val != 2 || node.Right.Left.Val != 3 || node.Right.Right != nil {
		t.Fatal("sparse arrays must assign children to non-null parents, not heap indices")
	}
}

func TestBinaryTreeRejectsMalformedArrays(t *testing.T) {
	kind := reflect.TypeOf((*testTreeNode)(nil))
	for _, raw := range []string{`[null,1]`, `[1,null,null,2]`, `[1,"2"]`, `[1,1.5]`, `[1,[2]]`, `[true]`} {
		if _, err := decode(json.RawMessage(raw), kind); err == nil {
			t.Fatalf("invalid tree accepted: %s", raw)
		}
	}
}

func TestBinaryTreeDecodingCreatesIndependentNodes(t *testing.T) {
	kind := reflect.TypeOf((*testTreeNode)(nil))
	first, err := decode(json.RawMessage(`[1,2,2]`), kind)
	if err != nil {
		t.Fatal(err)
	}
	second, err := decode(json.RawMessage(`[1,2,2]`), kind)
	if err != nil {
		t.Fatal(err)
	}
	root := first.Interface().(*testTreeNode)
	root.Left.Val = 99
	if root.Right.Val != 2 || second.Interface().(*testTreeNode).Left.Val != 2 {
		t.Fatal("tree nodes were shared across branches or invocations")
	}
}

func TestBinaryTreeComparisonPreservesStructure(t *testing.T) {
	kind := reflect.TypeOf((*testTreeNode)(nil))
	for _, check := range []struct {
		got, want string
		equal     bool
	}{
		{`[1,2,3]`, `[1,2,3,null,null,null,null]`, true},
		{`[1,2]`, `[1,null,2]`, false},
		{`[1,2,3]`, `[1,3,2]`, false},
		{`[]`, `null`, true},
	} {
		got, err := decode(json.RawMessage(check.got), kind)
		if err != nil {
			t.Fatal(err)
		}
		want, err := decode(json.RawMessage(check.want), kind)
		if err != nil {
			t.Fatal(err)
		}
		equal, err := matches(got, want, "exact")
		if err != nil || equal != check.equal {
			t.Fatalf("%s vs %s: equal=%v, want %v (%v)", check.got, check.want, equal, check.equal, err)
		}
	}
}

func TestBinaryTreeRejectsCyclesAndSharedNodes(t *testing.T) {
	for _, cycle := range []bool{true, false} {
		root := &testTreeNode{Val: 1}
		if cycle {
			root.Left = root
		} else {
			child := &testTreeNode{Val: 2}
			root.Left, root.Right = child, child
		}
		if _, err := matches(reflect.ValueOf(root), reflect.ValueOf((*testTreeNode)(nil)), "exact"); err == nil {
			t.Fatal("a cycle or shared child must not be accepted as a tree")
		}
	}
}

func TestCombinationsRejectsFlatResults(t *testing.T) {
	if _, err := matches(reflect.ValueOf([]int{}), reflect.ValueOf([]int{}), "combinations"); err == nil {
		t.Fatal("flat combinations accepted")
	}
}

func TestRunCombinations(t *testing.T) {
	path := fixture(t, `[{"id":1,"input":{"values":[2,3,7]},"expected":[[2,2,3],[7]]}]`)
	Run(t, path, func(values []int) [][]int { return [][]int{{7}, {3, 2, 2}} }, []string{"values"}, "combinations")
}
