# Testing and benchmarks

One shared Go runner loads cases, decodes typed arguments, invokes the solution, and compares outputs. A command reads Go signatures and generates tiny adapters automatically. Benchmarking is not implemented yet.

## Running tests

From the repository root:

```sh
# One implementation
go run ./cmd/leetcode test problems/algorithms/0001-two-sum/solutions/go/hashmap

# All implementations of one problem
go run ./cmd/leetcode test problems/algorithms/0001-two-sum

# Longest Substring Without Repeating Characters
go run ./cmd/leetcode test problems/algorithms/0003-longest-substring-without-repeating-characters

# Valid Parentheses
go run ./cmd/leetcode test problems/algorithms/0020-valid-parentheses

# Every Go solution
go run ./cmd/leetcode test

# One case of one solution
go run ./cmd/leetcode test --run '^TestSolution/case_1$' problems/algorithms/0001-two-sum/solutions/go/hashmap
```

Every test run regenerates adapters before calling `go test -v -count=1`, with a 30-second timeout per package. Each numeric case appears as a named subtest. Failed cases and generation errors cause the command to exit unsuccessfully.

`go run ./cmd/leetcode generate` only generates adapters. After generation, ordinary `go test` commands also work, but they do not refresh adapters if signatures or comparison settings change. Fresh clones need generation first.

Before each commit, the hook generates adapters and runs all Go and Python tests in a temporary copy of the staged repository. An unstaged fix cannot hide a failing staged solution or fixture. Failures block the commit before the README is regenerated; tests do not modify your working solution files.

## How matching works

Keep implementations in `solutions/go/<approach>/solution.go`. The command scans source files in that package with Go's parser and finds a function whose named parameters match the first case's input keys. It ignores methods and test files. Parameters are passed in signature order; JSON key order is irrelevant.

For `func twoSum(nums []int, target int) []int`, inputs must contain `nums` and `target`. The shared runner uses the function's compiled types to decode each value and its expected result. Missing, extra, or incorrectly typed arguments fail rather than receiving default values.

If several functions match, choose one explicitly:

```sh
go run ./cmd/leetcode test --function twoSum problems/algorithms/0001-two-sum/solutions/go/hashmap
```

The current runner supports top-level functions with fixed named parameters and exactly one returned value. Integers, strings, booleans, slices, arrays, maps, and other types accepted by Go's JSON decoder work directly. Generic functions, variadic functions, methods, multiple-return functions, and LeetCode-specific tree/list encodings require additional adapters or conversion support; the runner does not guess them.

Generated `zz_generated_test.go` files are ignored by Git. They contain a function reference, parameter names, fixture path, and comparison mode; they do not contain solution code or custom fixture parsing. Do not maintain them manually. The generator only overwrites files bearing its generated header and never edits `solution.go` or handwritten tests.

## Shared fixtures

Keep one `testcases.json` beside each problem's `problem.md`. All implementations use the same cases, including custom tests. Format one case per line with inline arrays:

```json
[
  {"id": 1, "input": {"nums": [2, 7, 11, 15], "target": 9}, "expected": [0, 1], "remark": ""},
  {"id": 2, "input": {"nums": [3, 2, 4], "target": 6}, "expected": [1, 2], "remark": ""}
]
```

- `id`: a positive integer unique within the file. LeetCode cases keep their example numbers; append custom cases using the next unused number.
- `input`: named function arguments. In Two Sum, `target` is the requested sum.
- `expected`: the returned result. In Two Sum, this is the pair of indices.
- `remark`: blank by default; use it to explain a custom case. It may be omitted.

Appending a case automatically includes it in the next run. Missing or empty fixtures, malformed JSON, duplicate IDs, missing required values, and unknown case fields fail the suite. Each invocation decodes fresh arguments, so mutable inputs are not reused between cases. Panics are reported as failed subtests.

## Comparison rules

Exact typed equality is the default. For unordered output, add one line to the problem's `problem.md`:

```text
Comparison: unordered
```

Two Sum uses this setting to accept either index order. Unordered comparison ignores only the outer array order, preserves duplicate counts, and preserves order inside nested arrays. Null and empty collections remain distinct. Unknown comparison settings fail generation. Floating-point tolerance and property validators are not implemented yet.

The runner matches the stored expected result, so cases with multiple valid answers beyond ordering will need a validator when support is added. Local test success does not establish LeetCode acceptance.

## Performance measurements

Verify correctness first, then repeat measurements per case. Exclude compilation, fixture parsing, input preparation, and comparisons from algorithm timing. Rebuild mutable inputs outside the measured interval and include the solution's own allocations.

Report median time per execution, allocated bytes per execution, and allocation count. Estimate suite time by summing per-case median times and identify the slowest case. Allocated bytes are not peak memory usage. An optional separate compiled-process measurement can report peak resident memory, including runtime and harness overhead.

Compare approaches with the same cases and environment, running benchmarks sequentially. Record the commit, fixtures, toolchain, OS, CPU, repetition policy, and timestamp with saved results. These are local comparisons, not LeetCode percentiles; avoid a combined score.

Keep generated reports under ignored `.local/`. Correctness checks can run in future CI; benchmarks should initially run on demand.
