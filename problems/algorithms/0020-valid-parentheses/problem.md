# 20. Valid Parentheses

https://leetcode.com/problems/valid-parentheses/

Difficulty: Easy
Topics: String, Stack

The statement below is paraphrased.

## Statement

Decide whether the bracket string `s` is balanced. Each closing bracket must match the most recent opening bracket that has not yet been matched, using the same bracket type. No opening or closing bracket may remain unmatched. The permitted types are round, square, and curly brackets.

## Examples

| Case | `s` | Expected result |
| --- | --- | --- |
| 1 | `"()"` | `true` |
| 2 | `"()[]{}"` | `true` |
| 3 | `"(]"` | `false` |
| 4 | `"([])"` | `true` |
| 5 | `"([)]"` | `false` |

The same examples are stored in [testcases.json](testcases.json).

## Constraints

- String size: 1 through 10,000 characters, inclusive.
- Every character belongs to `()[]{}`.

## Output rules

Return a boolean and compare it exactly. Matching counts alone are insufficient: bracket types and nesting order must also match.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/stack](solutions/go/stack/) |  |  |  |  |  |
