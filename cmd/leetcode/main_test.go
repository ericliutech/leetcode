package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestDiscoverUsesParameterNamesAndGroupedParameters(t *testing.T) {
	directory := t.TempDir()
	write(t, filepath.Join(directory, "solution.go"), `package solution
func add(right, left int) (answer int) { return right + left }
func helper(value int) int { return value }
`)
	packageName, function, err := discoverEntry(directory, []string{"left", "right"}, "")
	if err != nil {
		t.Fatal(err)
	}
	if packageName != "solution" || function.name != "add" || strings.Join(function.parameters, ",") != "right,left" {
		t.Fatalf("wrong signature discovery: %s %#v", packageName, function)
	}
}

func TestAmbiguousEntryRequiresSelection(t *testing.T) {
	directory := t.TempDir()
	write(t, filepath.Join(directory, "solution.go"), `package solution
func fast(s string) int { return len(s) }
func slow(s string) int { return len(s) }
`)
	if _, _, err := discoverEntry(directory, []string{"s"}, ""); err == nil {
		t.Fatal("ambiguous functions were accepted")
	}
	if _, selected, err := discoverEntry(directory, []string{"s"}, "fast"); err != nil || selected.name != "fast" {
		t.Fatalf("explicit selection failed: %v", err)
	}
}

func temporaryRepository(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	write(t, filepath.Join(root, "go.mod"), "module github.com/ericliutech/leetcode\n\ngo 1.26.1\n")
	problem := filepath.Join(root, "problems/algorithms/0099-example")
	directory := filepath.Join(problem, "solutions/go/example")
	write(t, filepath.Join(problem, "problem.md"), "# 99. Example\n")
	write(t, filepath.Join(problem, "testcases.json"), `[{"id":1,"input":{"right":4,"left":3},"expected":7}]`)
	write(t, filepath.Join(directory, "solution.go"), "package solution\nfunc sum(left, right int) int { return left + right }\n")
	return root, directory
}

func within(t *testing.T, root string) {
	t.Helper()
	original, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(original); err != nil {
			t.Error(err)
		}
	})
}

func TestGenerationProtectsHandwrittenFilesAndSource(t *testing.T) {
	root, directory := temporaryRepository(t)
	within(t, root)
	source, _ := os.ReadFile(filepath.Join(directory, "solution.go"))
	path := filepath.Join(directory, generatedName)
	write(t, path, "package solution\n// Handwritten test.\n")
	if err := run([]string{"generate"}); err == nil || !strings.Contains(err.Error(), "handwritten") {
		t.Fatalf("handwritten file protection failed: %v", err)
	}
	if content, _ := os.ReadFile(path); string(content) != "package solution\n// Handwritten test.\n" {
		t.Fatal("handwritten file was changed")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := run([]string{"generate"}); err != nil {
		t.Fatal(err)
	}
	if content, _ := os.ReadFile(filepath.Join(directory, "solution.go")); string(content) != string(source) {
		t.Fatal("solution source changed")
	}
	adapter, _ := os.ReadFile(path)
	if !strings.Contains(string(adapter), `sum, []string{"left", "right"}`) {
		t.Fatalf("missing typed adapter: %s", adapter)
	}
}

func TestUnknownComparisonIsRejected(t *testing.T) {
	root, directory := temporaryRepository(t)
	problem := filepath.Dir(filepath.Dir(filepath.Dir(directory)))
	for _, comparison := range []string{"typo", ""} {
		write(t, filepath.Join(problem, "problem.md"), "# 99. Example\nComparison: "+comparison+"\n")
		if _, _, err := prepare(root, directory, ""); err == nil {
			t.Fatal("invalid comparison was accepted")
		}
	}
}

func TestGeneratedAdapterCompilesAndRunsANewSignature(t *testing.T) {
	root, directory := temporaryRepository(t)
	shared, err := os.ReadFile(filepath.Join("..", "..", "internal", "judge", "judge.go"))
	if err != nil {
		t.Fatal(err)
	}
	write(t, filepath.Join(root, "internal/judge/judge.go"), string(shared))
	within(t, root)
	if err := run([]string{"test", directory}); err != nil {
		t.Fatalf("new signature did not compile and pass: %v", err)
	}
	problem := filepath.Dir(filepath.Dir(filepath.Dir(directory)))
	write(t, filepath.Join(problem, "testcases.json"), `[{"id":1,"input":{"right":4,"left":3},"expected":99}]`)
	if err := run([]string{"test", directory}); err == nil {
		t.Fatal("incorrect expected output did not fail the command")
	}
}
