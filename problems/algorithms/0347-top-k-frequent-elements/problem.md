# 347. Top K Frequent Elements

https://leetcode.com/problems/top-k-frequent-elements/

Difficulty: Medium
Topics: Array, Hash table, Heap
Comparison: unordered

The statement below is paraphrased.

## Statement

Given `nums` and `k`, return the `k` distinct values with the highest occurrence counts. The selected set is guaranteed to be unique, and its values may be returned in any order.

## Examples

| Case | `nums` | `k` | Expected values |
| --- | --- | --- | --- |
| 1 | `[1, 1, 1, 2, 2, 3]` | 2 | `[1, 2]` |
| 2 | `[1]` | 1 | `[1]` |
| 3 | `[1, 2, 1, 2, 1, 2, 3, 1, 3, 2]` | 2 | `[1, 2]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Array size: 1 through 100,000 elements, inclusive.
- Each value is between -10,000 and 10,000, inclusive.
- `k` is at least 1 and no larger than the number of distinct values.
- The highest-frequency set of size `k` is unambiguous.

## Output rules

Return values, not their frequencies or indices. Compare without regard to order, preserving multiplicity so duplicate returned values cannot replace a missing answer. Tied frequencies inside the selected set are allowed.

## Follow-up

Seek a time bound strictly better than O(n log n). This implementation has an O(n log n) worst-case upper bound and does not guarantee that follow-up requirement.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/max-heap](solutions/go/max-heap/) | Select the k most frequent distinct values | Count frequencies, heapify value/count pairs, then extract the maximum k times. | O(n + k log m) expected | O(m) |  |

`n` is the input length; `m` is the number of distinct values. Counting takes expected O(n), bottom-up heap construction takes O(m), and k removals take O(k log m). Since m ≤ n, the total simplifies to O(n + k log m). The map and heap use O(m) space; k ≤ m. Heapify is linear, but the repeated removals mean the full algorithm is not generally linear. The solution does not modify `nums`.
