# 39. Combination Sum

https://leetcode.com/problems/combination-sum/

Difficulty: Medium
Topics: Array, Backtracking
Comparison: combinations

The statement below is paraphrased.

## Statement

Find every distinct combination of values from `candidates` that adds up to `target`. Each candidate may be reused any number of times. Different orderings of the same values count as one combination.

## Examples

| Case | `candidates` | `target` | Expected combinations |
| --- | --- | --- | --- |
| 1 | `[2, 3, 6, 7]` | `7` | `[[2, 2, 3], [7]]` |
| 2 | `[2, 3, 5]` | `8` | `[[2, 2, 2, 2], [2, 3, 3], [3, 5]]` |
| 3 | `[2]` | `1` | `[]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Between 1 and 30 distinct candidates, inclusive.
- Candidate values are between 2 and 40, inclusive.
- `target` is between 1 and 40, inclusive.
- Each input has fewer than 150 valid combinations.

## Output rules

Compare combinations regardless of outer or inner order, preserving value frequencies and rejecting duplicate combinations. No-match fixtures use `null` for this implementation's nil result slice; the runner distinguishes nil from an allocated empty slice.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/backtracking](solutions/go/backtracking/) | Enumerate combinations with reusable choices | Backtrack: choose, recurse from the same index for reuse, then undo; clone successful paths. Sorting enables early pruning. | Exponential | O(d) auxiliary; output excluded |  |

`d = target / min(candidates)` bounds path length. Sorting adds O(n log n) and mutates the input; pruning reduces work but search remains exponential. A conservative total bound is O(n log n + (n + 1)^(d + 1)), including path copies. Nondecreasing indices prevent duplicate permutations; cloning keeps saved paths independent.
