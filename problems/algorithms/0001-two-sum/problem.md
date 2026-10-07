# 1. Two Sum

https://leetcode.com/problems/two-sum/

Difficulty: Easy
Topics: Array, Hash table
Comparison: unordered

The statement below is paraphrased.

## Statement

Given the integer array `nums` and integer `target`, identify two different positions whose values sum to the target. Return their zero-based indices. The inputs guarantee a single valid pair; either index order is acceptable.

## Examples

| Case | `nums` | `target` | Expected indices |
| --- | --- | ---: | --- |
| 1 | `[2, 7, 11, 15]` | 9 | `[0, 1]` |
| 2 | `[3, 2, 4]` | 6 | `[1, 2]` |
| 3 | `[3, 3]` | 6 | `[0, 1]` |

The same examples are stored in [testcases.json](testcases.json).

## Constraints

- Array size: 2 through 10,000 elements, inclusive.
- Each element and the target are within -1,000,000,000 through 1,000,000,000.
- Exactly one pair of distinct indices satisfies the target.

## Output rules

Return indices, not values. Compare the two indices without regard to order, preserving multiplicity; selecting one position twice is invalid.

## Follow-up

Consider an approach with time complexity better than quadratic.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/hashmap](solutions/go/hashmap/) | Pair sum; fast complement lookup | Check `target-num` among earlier values, then store `num → index`. | O(n) expected | O(n) |  |
