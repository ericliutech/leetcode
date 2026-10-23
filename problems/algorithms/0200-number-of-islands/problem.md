# 200. Number of Islands

https://leetcode.com/problems/number-of-islands/

Difficulty: Medium
Topics: Array, Depth-first search, Breadth-first search, Union find, Matrix

The statement below is paraphrased.

## Statement

Count the islands in a rectangular grid containing land (`'1'`) and water (`'0'`). Land cells belong to the same island when connected horizontally or vertically. Diagonal contact does not connect islands. Water surrounds the grid.

## Examples

| Case | `grid` (one string per row) | Expected islands |
| --- | --- | --- |
| 1 | `["11110", "11010", "11000", "00000"]` | `1` |
| 2 | `["11000", "11000", "00100", "00011"]` | `3` |

These examples and additional edge cases are stored in [testcases.json](testcases.json). The Go function accepts `[][]byte`; fixtures encode cells as numeric ASCII bytes: `48` for `'0'`, `49` for `'1'`. This uses the shared runner's existing JSON decoding without changing the submission signature.

## Constraints

- The grid has between 1 and 300 rows and between 1 and 300 columns.
- Every row has the same length.
- Each cell is either `'0'` or `'1'`.

## Output rules

Compare the returned island count exactly. The implementation mutates the input grid to mark visited land; each test invocation receives fresh input.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/bfs](solutions/go/bfs/) | Count connected components in a grid | Scan for unvisited land; count it and flood-fill its island with a queue, marking cells when enqueued. | O(m × n) | O(m × n) worst case |  |

`m` and `n` are the row and column counts. Marking cells when enqueuing prevents duplicate visits; either BFS or DFS can use that rule. The grid doubles as the visited bitmap.

In this implementation, `clear(queue)` runs after the previous traversal has drained the slice, so it is a no-op. Dropping the slice prefix avoids retaining references to visited objects because queue elements contain only integers; the slice can still retain its backing allocation. The supplied solution is preserved unchanged apart from its package declaration.
