# 560. Subarray Sum Equals K

https://leetcode.com/problems/subarray-sum-equals-k/

Difficulty: Medium
Topics: Array, Hash table, Prefix sum

The statement below is paraphrased.

## Statement

Count the non-empty contiguous portions of `nums` whose elements add up to `k`. Count different index ranges separately, even when their values match.

## Examples

| Case | `nums` | `k` | Expected count |
| --- | --- | --- | --- |
| 1 | `[1, 1, 1]` | `2` | `2` |
| 2 | `[1, 2, 3]` | `3` | `2` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Array length: 1 through 20,000, inclusive.
- Each element is between -1,000 and 1,000, inclusive.
- `k` is between -10,000,000 and 10,000,000, inclusive.

## Output rules

Compare the returned count exactly. Overlapping subarrays count; empty subarrays do not.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/prefix-sum-hashmap](solutions/go/prefix-sum-hashmap/) | Count contiguous ranges with an exact sum; negatives prevent standard sliding-window reasoning | For each prefix sum S, count earlier prefixes equal to S - k using a frequency map. | O(n) expected | O(n) | Started with O(n²); understood this approach with ChatGPT help. |

Seed `prefix[0] = 1` to count ranges starting at index 0. Look up before incrementing the current prefix frequency to exclude empty ranges when `k = 0`. `n` is the array length.
