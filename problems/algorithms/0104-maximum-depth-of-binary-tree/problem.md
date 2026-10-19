# 104. Maximum Depth of Binary Tree

https://leetcode.com/problems/maximum-depth-of-binary-tree/

Difficulty: Easy
Topics: Tree, Depth-first search, Binary tree

The statement below is paraphrased.

## Statement

Find the largest number of nodes on a path from `root` to a leaf. Count the root as depth 1; an empty tree has depth 0.

## Examples

| Case | `root` | Expected depth |
| --- | --- | --- |
| 1 | `[3, 9, 20, null, null, 15, 7]` | 3 |
| 2 | `[1, null, 2]` | 2 |

These examples and additional edge cases are stored in [testcases.json](testcases.json). Trees use breadth-first arrays with `null` for missing children.

## Constraints

- Tree size: 0 through 10,000 nodes, inclusive.
- Each node's value is between -100 and 100, inclusive.

## Output rules

Return an integer and compare it exactly. Depth counts nodes, not edges.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/iterative-dfs](solutions/go/iterative-dfs/) | Track depth along root-to-leaf paths | Store node/depth pairs on a stack, pop each node, update the maximum, and push its children at the next depth. | O(n) | O(h) upper bound; O(n) worst case |  |
| [go/recursive-dfs](solutions/go/recursive-dfs/) | Tree depth combines the depths of its subtrees | Return 0 for nil; otherwise return 1 plus the larger child depth. | O(n) | O(h) call stack; O(n) worst case |  |

`n` is the node count; `h` is the maximum depth. Both solutions visit every node once and leave the tree unchanged. Balanced trees have h = O(log n); a chain can have h = n.

The iterative stack holds pending branches and records depth explicitly. On a single-child chain it contains only one entry, whereas recursion still uses O(h) call frames.
