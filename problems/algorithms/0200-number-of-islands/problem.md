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

Compare the returned island count exactly. The in-place BFS and recursive DFS mutate the grid; the visited-bitmap BFS preserves it. Each test invocation receives fresh input.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/bfs-in-place](solutions/go/bfs-in-place/) | Count connected components; input may be modified | Count each unvisited island and flood-fill it, marking land directly in the grid when enqueued. | O(m × n) | O(m × n) |  |
| [go/bfs-bitmap](solutions/go/bfs-bitmap/) | Count connected components; preserve the input | Track visited cells separately and flood-fill each island with a queue. Mark neighbors when enqueued; only expand from land. | O(m × n) | O(m × n) |  |
| [go/dfs-recursive](solutions/go/dfs-recursive/) | Count connected components using recursive flood-fill | Count each unvisited island; mark land before recursively exploring its four neighbors. | O(m × n) | O(m × n) worst-case call stack |  |

`m` and `n` are the row and column counts. All versions expand each land cell once and check at most four neighbors, giving O(m × n) total time. BFS marks cells when enqueuing to prevent duplicates; the bitmap version also enqueues neighboring water cells, but only expands from land.

Both BFS versions preallocate a buffer with capacity `m*n`, so auxiliary space is O(m × n) for every input, including an all-water grid. The in-place version marks visited land in the grid; the bitmap version preserves the grid by allocating an additional O(m × n) boolean matrix.

Each island starts with `queue := buffer`. Advancing `queue` with `queue[1:]` leaves the original buffer header unchanged, so the next island reuses its full capacity from the beginning. The capacity covers every possible enqueue in one traversal. No clearing is needed: entries contain only integers and appends overwrite old values.

Recursive DFS uses the call stack instead of a queue. Marking land as `'x'` before exploring prevents repeat expansion. Visited tracking needs O(1) extra space, but a traversal can reach O(m × n) recursion depth, so total auxiliary space is not O(1).
