# 15. 3Sum

https://leetcode.com/problems/3sum/

Difficulty: Medium
Topics: Array, Two pointers, Sorting
Comparison: unordered

The statement below is paraphrased.

## Statement

Find every distinct combination of three values from `nums` whose sum is zero. Each combination must use three different array positions. Return the values, not their indices, and include each unique triplet only once.

## Examples

| Case | `nums` | Expected triplets |
| --- | --- | --- |
| 1 | `[-1, 0, 1, 2, -1, -4]` | `[[-1, -1, 2], [-1, 0, 1]]` |
| 2 | `[0, 1, 1]` | `[]` |
| 3 | `[0, 0, 0]` | `[[0, 0, 0]]` |

These cases are stored in [testcases.json](testcases.json).

## Constraints

- Array size: 3 through 3,000 elements, inclusive.
- Each value is between -100,000 and 100,000, inclusive.

## Output rules

LeetCode accepts any ordering of the triplets and of the values within each triplet. Duplicate triplets are not allowed.

The current shared runner ignores outer array order only. These fixtures list each triplet's values in ascending order; arbitrary inner ordering will require additional comparison support.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/two-pointer](solutions/go/two-pointer/) | Zero-sum triplets; sorting guides pair search | Sort, fix one value, move two pointers by the sum, and skip duplicates. | O(n²) | O(log n) sorting stack | A hash-based pair search complicates duplicate handling; sorting lets me skip duplicates directly. |

`n` is the input length. The pointer scan uses O(1) auxiliary space. Sorting modifies `nums` in place.
