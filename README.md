# LeetCode solutions

My LeetCode problems, solutions, and notes for future review.

<!-- generated:progress:start -->

## Progress

| Measure | Count |
| --- | ---: |
| Documented problems | 3 |
| Solution entries | 3 |

Counts reflect documentation, not correctness or LeetCode acceptance.

## Problem index

| ID | Problem | Difficulty | Topics | Languages |
| --- | --- | --- | --- | --- |
| 1 | [Two Sum](problems/algorithms/0001-two-sum/problem.md) | Easy | Array, Hash table | go |
| 3 | [Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | Medium | Hash table, String, Sliding window | go |
| 20 | [Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | Easy | String, Stack | go |

## Solution notes

| Problem | Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- | --- |
| [1. Two Sum](problems/algorithms/0001-two-sum/problem.md) | [go/hashmap](problems/algorithms/0001-two-sum/solutions/go/hashmap/) | Pair sum; fast complement lookup | Check `target-num` among earlier values, then store `num → index`. | O(n) expected | O(n) |  |
| [3. Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | [go/sliding-window](problems/algorithms/0003-longest-substring-without-repeating-characters/solutions/go/sliding-window/) | Longest contiguous segment with no duplicates | Track last-seen indices; move `left` past a duplicate inside the window and update the maximum length. | O(n) | O(1), fixed 256-entry array | Tried recursive DP; a sliding window tracks uniqueness directly. |
| [20. Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | [go/stack](problems/algorithms/0020-valid-parentheses/solutions/go/stack/) | Nested pairs; last opened must close first | Push opening brackets. For each closing bracket, require a matching stack top, then pop. Accept only if the stack ends empty. | Θ(n) worst case | O(n) |  |

<!-- generated:progress:end -->

## Using this repository

Each problem has a `problem.md` containing its statement, facts, and solution notes; a compact `testcases.json`; and its implementation files.

- [Conventions](docs/conventions.md): adding problems and updating the generated overview.
- [Testing and benchmarks](docs/testing-and-benchmarks.md): fixture format and measurement policy.

Edit `problem.md` to record your approach, trigger, key idea, complexity, and mistakes. The progress, index, and solution notes above are generated from those files.

Run `python3 scripts/consolidate.py` to refresh the overview, or let the pre-commit hook refresh it from staged documentation. Use `python3 scripts/consolidate.py --check` to check that it is current.

The pre-commit hook tests a temporary copy of all staged files before updating the README. Solution cases, shared runner tests, and repository automation tests must pass for the commit to proceed.

Go solutions share the root `go.mod`. Run `go run ./cmd/leetcode test` to generate adapters and check every Go implementation against its `testcases.json`, or see the [individual commands](docs/testing-and-benchmarks.md#running-tests). Benchmarking and the CI workflow are not implemented yet. Solution entries do not imply local verification or LeetCode acceptance.

Solution source files are maintained by the repository owner. Automated tooling must not edit, format, move, or rename them without explicit permission.
