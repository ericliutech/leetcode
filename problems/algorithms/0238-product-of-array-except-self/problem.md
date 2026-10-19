# 238. Product of Array Except Self

https://leetcode.com/problems/product-of-array-except-self/

Difficulty: Medium
Topics: Array, Prefix sum

The statement below is paraphrased.

## Statement

For each position in `nums`, compute the product of every value at the other positions. Return those products in the same index order. Use O(n) time and do not use division.

## Examples

| Case | `nums` | Expected products |
| --- | --- | --- |
| 1 | `[1, 2, 3, 4]` | `[24, 12, 8, 6]` |
| 2 | `[-1, 1, 0, -3, 3]` | `[0, 0, 9, 0, 0]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Array size: 2 through 100,000 elements, inclusive.
- Each value is between -30 and 30, inclusive.
- Prefix products, suffix products, and each returned value fit in a signed 32-bit integer.

## Output rules

Compare the result exactly, preserving index order. Exclude only the current position, not all occurrences of its value.

## Follow-up

Use O(1) auxiliary space, excluding the returned array. This implementation meets that bound.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/prefix-suffix](solutions/go/prefix-suffix/) | Exclude one index from a product; division is disallowed | Store exclusive prefix products in the result, then multiply by a running suffix product while scanning backward. | O(n) | O(1) auxiliary; output excluded |  |

`n` is the array length. The forward pass stores the product strictly to the left of each index. The backward pass compounds values strictly to its right. Their product excludes exactly the current element and naturally handles zeros.

The prefix at index 0 and the suffix after the last index are both 1. The last result entry already contains its complete prefix product, so the backward pass begins at n - 2. The solution leaves `nums` unchanged and uses the returned array as prefix storage; including that array, total space is O(n).
