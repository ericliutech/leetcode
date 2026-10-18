# 347. Top K Frequent Elements

https://leetcode.com/problems/top-k-frequent-elements/

Difficulty: Medium
Topics: Array, Hash table, Heap, Bucket sort
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

Seek a time bound strictly better than O(n log n). Neither heap implementation guarantees that bound for every allowed value of k; the bucket approach achieves expected O(n).

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/max-heap](solutions/go/max-heap/) | Select the k most frequent distinct values | Count frequencies, heapify value/count pairs, then extract the maximum k times. | O(n + k log m) expected | O(m) |  |
| [go/min-heap](solutions/go/min-heap/) | Top k; retain only k candidates | Count frequencies, keep a size-k min-heap, and replace its root when a higher count appears. | O(n + m log k) expected | O(m) total; O(k) heap |  |
| [go/bucket-sort](solutions/go/bucket-sort/) | Frequencies are bounded by the input length | Count frequencies and track the maximum; group values by count, then scan buckets downward until k values are collected. | O(n) expected | O(m + F), worst-case O(n) |  |

`n` is the input length; `m` is the number of distinct values; `F` is the maximum frequency. Here k ≤ m ≤ n and F ≤ n. All three approaches count frequencies in expected O(n) time, use an O(m) frequency map, and leave `nums` unchanged.

- Max-heap: bottom-up heap construction takes O(m), then k removals take O(k log m). Heap storage is O(m).
- Min-heap: examine m entries, with each heap insertion or replacement costing at most O(log k), then copy k values into the result. Heap storage is O(k), but total space remains O(m) because of the frequency map.
- Bucket sort: allocate F + 1 buckets once, distribute m distinct values, and scan downward to collect k results. Total expected time is O(n + m + F + k) = O(n). Space is O(m + F); a large frequency can create many empty buckets even when m is small.

A smaller heap does not guarantee a faster run: the max-heap version builds its heap in linear time. For fixed k, both heap approaches have an O(n) expected bound. Bucket allocation and scanning also have costs, so its stronger general time bound need not yield a faster measured run. Similar LeetCode scores do not establish which implementation is faster on other inputs.
