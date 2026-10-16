# 33. Search in Rotated Sorted Array

https://leetcode.com/problems/search-in-rotated-sorted-array/

Difficulty: Medium
Topics: Array, Binary search

The statement below is paraphrased.

## Statement

An array of distinct integers was sorted in increasing order and may then have been rotated by moving a prefix to the end. Given this array as `nums` and an integer `target`, find the target's zero-based position in the supplied array. Return `-1` if it is absent. The search must take O(log n) time.

## Examples

| Case | `nums` | `target` | Expected index |
| --- | --- | --- | --- |
| 1 | `[4, 5, 6, 7, 0, 1, 2]` | 0 | 4 |
| 2 | `[4, 5, 6, 7, 0, 1, 2]` | 3 | -1 |
| 3 | `[1]` | 0 | -1 |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Array size: 1 through 5,000 elements, inclusive.
- Array elements and the target are within -10,000 through 10,000, inclusive.
- Every array value is unique.
- The array is increasing before the optional rotation.

## Output rules

Return an index in the supplied array, not in its unrotated form. Compare the integer result exactly.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/binary-search](solutions/go/binary-search/) | Rotated sorted array; O(log n) required | Compare target with the midpoint; use endpoint bounds to detect when rotation reverses the usual search direction. | O(log n) | O(1) | Tried recursion; iteration avoids extra stack space. |

`n` is the array length. Each iteration excludes the midpoint and discards one half. The solution does not modify `nums`.
