# 739. Daily Temperatures

https://leetcode.com/problems/daily-temperatures/

Difficulty: Medium
Topics: Array, Stack, Monotonic stack

The statement below is paraphrased.

## Statement

For each day, find the first later day with a strictly higher temperature. Return the distance in days, or 0 if no such day exists.

## Examples

| Case | `temperatures` | Expected waits |
| --- | --- | --- |
| 1 | `[73, 74, 75, 71, 69, 72, 76, 73]` | `[1, 1, 4, 2, 1, 1, 0, 0]` |
| 2 | `[30, 40, 50, 60]` | `[1, 1, 1, 0]` |
| 3 | `[30, 60, 90]` | `[1, 1, 0]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Between 1 and 100,000 daily readings, inclusive.
- Each temperature is between 30 and 100, inclusive.

## Output rules

Compare the result exactly, preserving day order. Equal temperatures do not count as warmer.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/monotonic-stack](solutions/go/monotonic-stack/) | Find the next strictly greater value to the right | Keep unresolved indices in a stack with non-increasing temperatures. A warmer day pops indices and records their waiting times. | O(n) | O(n) |  |

`n` is the number of days. Each index is pushed once and popped at most once, so the nested loops take O(n) total time. Equal temperatures stay on the stack. Unresolved days retain the default result of 0. The stack can hold O(n) indices; the output also takes O(n) space.
