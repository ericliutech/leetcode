# 206. Reverse Linked List

https://leetcode.com/problems/reverse-linked-list/

Difficulty: Easy
Topics: Linked list, Recursion

The statement below is paraphrased.

## Statement

You receive `head`, a pointer to the first node of a singly linked list. Reverse the links so that traversing the list visits the nodes in the opposite order, then return the new head. An empty list remains empty.

## Examples

| Case | Input list | Expected list |
| --- | --- | --- |
| 1 | `[1, 2, 3, 4, 5]` | `[5, 4, 3, 2, 1]` |
| 2 | `[1, 2]` | `[2, 1]` |
| 3 | `[]` | `[]` |

These cases are stored in [testcases.json](testcases.json). Arrays represent node values in traversal order; the solution itself receives and returns a node pointer.

## Constraints

- List size: 0 through 5,000 nodes, inclusive.
- Each node's value is between -5,000 and 5,000, inclusive.

## Output rules

Return the new head with the original values in reverse traversal order. The final node must point to nil. For an empty input, return nil, represented as `[]` in the fixtures. Output order matters.

## Follow-up

Implement both a loop-based approach and a recursive approach.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/iterative](solutions/go/iterative/) | Reverse links in place | Save the next node, reverse the current link, then advance. | O(n) | O(1) | |
| [go/recursion](solutions/go/recursion/) | Reverse the remaining list first | Reverse the suffix, point the next node back to the current node, then clear the current link. This is equivalent to using a stack iteratively. | O(n) | O(n), call stack | |
