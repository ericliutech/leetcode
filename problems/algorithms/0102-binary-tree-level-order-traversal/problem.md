# 102. Binary Tree Level Order Traversal

https://leetcode.com/problems/binary-tree-level-order-traversal/

Difficulty: Medium
Topics: Tree, Breadth-first search, Binary tree

The statement below is paraphrased.

## Statement

Starting at `root`, collect the values in each depth of a binary tree. Return one array per depth, beginning with the root's depth and listing nodes from left to right within each array. An empty tree produces an empty result.

## Examples

| Case | `root` | Expected levels |
| --- | --- | --- |
| 1 | `[3, 9, 20, null, null, 15, 7]` | `[[3], [9, 20], [15, 7]]` |
| 2 | `[1]` | `[[1]]` |
| 3 | `[]` | `[]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json). Input arrays encode nodes in breadth-first order, with `null` for missing children; trailing nulls may be omitted.

## Constraints

- Tree size: 0 through 2,000 nodes, inclusive.
- Each node's value is between -1,000 and 1,000, inclusive.

## Output rules

Compare both level order and the left-to-right order within each level exactly. Return `[]` for an empty tree.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/bfs](solutions/go/bfs/) | Visit a tree level by level | Process the current queue length, enqueue left then right children, and clear and drop the processed prefix. | O(n) | O(w) auxiliary queue |  |

`n` is the number of nodes; `w` is the maximum level width. Returned results are excluded from auxiliary space. Go's range loop captures the initial slice length, so appended children are processed in the next level. Clearing consumed queue entries removes stale node references.
