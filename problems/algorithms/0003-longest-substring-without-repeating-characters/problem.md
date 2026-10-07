# 3. Longest Substring Without Repeating Characters

https://leetcode.com/problems/longest-substring-without-repeating-characters/

Difficulty: Medium
Topics: Hash table, String, Sliding window

The statement below is paraphrased.

## Statement

For the string `s`, return the greatest length of a contiguous segment in which each character appears at most once. A segment must retain adjacent characters from the input; skipping characters does not form a substring.

## Examples

| Case | `s` | Expected length | One qualifying segment |
| --- | --- | ---: | --- |
| 1 | `"abcabcbb"` | 3 | `"abc"` |
| 2 | `"bbbbb"` | 1 | `"b"` |
| 3 | `"pwwkew"` | 3 | `"wke"` |

The same inputs and expected lengths are stored in [testcases.json](testcases.json).

## Constraints

- String size: 0 through 100,000 characters, inclusive.
- Inputs may contain English letters, digits, symbols, and spaces.

## Output rules

Return an integer length, not the segment itself. An empty input has length zero. Compare the returned integer exactly.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/sliding-window](solutions/go/sliding-window/) |  |  |  |  |  |
