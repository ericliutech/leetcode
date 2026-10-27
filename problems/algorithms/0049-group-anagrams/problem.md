# 49. Group Anagrams

https://leetcode.com/problems/group-anagrams/

Difficulty: Medium
Topics: Array, Hash table, String
Comparison: combinations

The statement below is paraphrased.

## Statement

Partition `strs` into groups whose strings contain exactly the same letters with the same frequencies. Include each input occurrence once, including repeated strings.

## Examples

| Case | `strs` | Expected groups |
| --- | --- | --- |
| 1 | `["eat", "tea", "tan", "ate", "nat", "bat"]` | `[["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]` |
| 2 | `[""]` | `[[""]]` |
| 3 | `["a"]` | `[["a"]]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Between 1 and 10,000 strings, inclusive.
- Each string has between 0 and 100 characters, inclusive.
- Characters are lowercase English letters only.

## Output rules

Ignore group order and string order within each group. Preserve duplicate string occurrences and group boundaries.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/frequency-array](solutions/go/frequency-array/) | Group strings by letter counts, regardless of character order | Use a 26-letter frequency array as the hash-map key; equal counts identify anagrams. | O(C + N) expected | O(N) |  |

`N` is the string count and `C` the total character count. Go arrays are comparable by value, so `[26]int` can be a map key directly. In languages without array keys, serializing the frequency map into an unambiguous string in a fixed character order is the next best alternative. Strings are referenced rather than copied.
