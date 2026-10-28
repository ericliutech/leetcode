# LeetCode solutions

My LeetCode problems, solutions, and notes for future review.

<!-- generated:progress:start -->

## Progress

| Measure | Count |
| --- | ---: |
| Documented problems | 16 |
| Solution entries | 22 |

Counts reflect documentation, not correctness or LeetCode acceptance.

## Problem index

| ID | Problem | Difficulty | Topics | Languages |
| --- | --- | --- | --- | --- |
| 1 | [Two Sum](problems/algorithms/0001-two-sum/problem.md) | Easy | Array, Hash table | go |
| 3 | [Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | Medium | Hash table, String, Sliding window | go |
| 15 | [3Sum](problems/algorithms/0015-3sum/problem.md) | Medium | Array, Two pointers, Sorting | go |
| 20 | [Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | Easy | String, Stack | go |
| 33 | [Search in Rotated Sorted Array](problems/algorithms/0033-search-in-rotated-sorted-array/problem.md) | Medium | Array, Binary search | go |
| 39 | [Combination Sum](problems/algorithms/0039-combination-sum/problem.md) | Medium | Array, Backtracking | go |
| 49 | [Group Anagrams](problems/algorithms/0049-group-anagrams/problem.md) | Medium | Array, Hash table, String | go |
| 56 | [Merge Intervals](problems/algorithms/0056-merge-intervals/problem.md) | Medium | Array, Sorting | go |
| 102 | [Binary Tree Level Order Traversal](problems/algorithms/0102-binary-tree-level-order-traversal/problem.md) | Medium | Tree, Breadth-first search, Binary tree | go |
| 104 | [Maximum Depth of Binary Tree](problems/algorithms/0104-maximum-depth-of-binary-tree/problem.md) | Easy | Tree, Depth-first search, Binary tree | go |
| 200 | [Number of Islands](problems/algorithms/0200-number-of-islands/problem.md) | Medium | Array, Depth-first search, Breadth-first search, Union find, Matrix | go |
| 206 | [Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | Easy | Linked list, Recursion | go |
| 238 | [Product of Array Except Self](problems/algorithms/0238-product-of-array-except-self/problem.md) | Medium | Array, Prefix sum | go |
| 347 | [Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | Medium | Array, Hash table, Heap, Bucket sort | go |
| 560 | [Subarray Sum Equals K](problems/algorithms/0560-subarray-sum-equals-k/problem.md) | Medium | Array, Hash table, Prefix sum | go |
| 739 | [Daily Temperatures](problems/algorithms/0739-daily-temperatures/problem.md) | Medium | Array, Stack, Monotonic stack | go |

## Solution notes

| Problem | Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- | --- |
| [1. Two Sum](problems/algorithms/0001-two-sum/problem.md) | [go/hashmap](problems/algorithms/0001-two-sum/solutions/go/hashmap/) | Pair sum; fast complement lookup | Check `target-num` among earlier values, then store `num → index`. | O(n) expected | O(n) |  |
| [3. Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | [go/sliding-window](problems/algorithms/0003-longest-substring-without-repeating-characters/solutions/go/sliding-window/) | Longest contiguous segment with no duplicates | Track last-seen indices; move `left` past a duplicate inside the window and update the maximum length. | O(n) | O(1), fixed 256-entry array | Tried recursive DP; a sliding window tracks uniqueness directly. |
| [15. 3Sum](problems/algorithms/0015-3sum/problem.md) | [go/two-pointer](problems/algorithms/0015-3sum/solutions/go/two-pointer/) | Zero-sum triplets; sorting guides pair search | Sort, fix one value, move two pointers by the sum, and skip duplicates. | O(n²) | O(log n) sorting stack | A hash-based pair search complicates duplicate handling; sorting lets me skip duplicates directly. |
| [20. Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | [go/stack](problems/algorithms/0020-valid-parentheses/solutions/go/stack/) | Nested pairs; last opened must close first | Push opening brackets. For each closing bracket, require a matching stack top, then pop. Accept only if the stack ends empty. | Θ(n) worst case | O(n) |  |
| [33. Search in Rotated Sorted Array](problems/algorithms/0033-search-in-rotated-sorted-array/problem.md) | [go/binary-search](problems/algorithms/0033-search-in-rotated-sorted-array/solutions/go/binary-search/) | Rotated sorted array; O(log n) required | Identify the sorted half, then use its bounds to decide which half can contain the target. | O(log n) | O(1) | Tried recursion; iteration avoids extra stack space. |
| [39. Combination Sum](problems/algorithms/0039-combination-sum/problem.md) | [go/backtracking](problems/algorithms/0039-combination-sum/solutions/go/backtracking/) | Enumerate combinations with reusable choices | Backtrack: choose, recurse from the same index for reuse, then undo; clone successful paths. Sorting enables early pruning. | Exponential | O(d) auxiliary; output excluded |  |
| [49. Group Anagrams](problems/algorithms/0049-group-anagrams/problem.md) | [go/frequency-array](problems/algorithms/0049-group-anagrams/solutions/go/frequency-array/) | Group strings by letter counts, regardless of character order | Use a 26-letter frequency array as the hash-map key; equal counts identify anagrams. | O(C + N) expected | O(N) |  |
| [56. Merge Intervals](problems/algorithms/0056-merge-intervals/problem.md) | [go/sort-and-merge](problems/algorithms/0056-merge-intervals/solutions/go/sort-and-merge/) | Overlapping ranges; sorted starts make merging local | Sort by start, merge into the last kept interval, and compact disjoint intervals into the same slice prefix. | O(n log n) | O(log n) sorting stack | Tried brute force O(n³); repeated slice splicing then made the merge pass O(n²). |
| [102. Binary Tree Level Order Traversal](problems/algorithms/0102-binary-tree-level-order-traversal/problem.md) | [go/bfs](problems/algorithms/0102-binary-tree-level-order-traversal/solutions/go/bfs/) | Visit a tree level by level | Process the current queue length, enqueue left then right children, and clear and drop the processed prefix. | O(n) | O(w) auxiliary queue |  |
| [104. Maximum Depth of Binary Tree](problems/algorithms/0104-maximum-depth-of-binary-tree/problem.md) | [go/iterative-dfs](problems/algorithms/0104-maximum-depth-of-binary-tree/solutions/go/iterative-dfs/) | Track depth along root-to-leaf paths | Store node/depth pairs on a stack, pop each node, update the maximum, and push its children at the next depth. | O(n) | O(h) upper bound; O(n) worst case |  |
| [104. Maximum Depth of Binary Tree](problems/algorithms/0104-maximum-depth-of-binary-tree/problem.md) | [go/recursive-dfs](problems/algorithms/0104-maximum-depth-of-binary-tree/solutions/go/recursive-dfs/) | Tree depth combines the depths of its subtrees | Return 0 for nil; otherwise return 1 plus the larger child depth. | O(n) | O(h) call stack; O(n) worst case |  |
| [200. Number of Islands](problems/algorithms/0200-number-of-islands/problem.md) | [go/bfs-in-place](problems/algorithms/0200-number-of-islands/solutions/go/bfs-in-place/) | Count connected components; input may be modified | Count each unvisited island and flood-fill it, marking land directly in the grid when enqueued. | O(m × n) | O(m × n) |  |
| [200. Number of Islands](problems/algorithms/0200-number-of-islands/problem.md) | [go/bfs-bitmap](problems/algorithms/0200-number-of-islands/solutions/go/bfs-bitmap/) | Count connected components; preserve the input | Track visited cells separately and flood-fill each island with a queue. Mark neighbors when enqueued; only expand from land. | O(m × n) | O(m × n) |  |
| [200. Number of Islands](problems/algorithms/0200-number-of-islands/problem.md) | [go/dfs-recursive](problems/algorithms/0200-number-of-islands/solutions/go/dfs-recursive/) | Count connected components using recursive flood-fill | Count each unvisited island; mark land before recursively exploring its four neighbors. | O(m × n) | O(m × n) worst-case call stack |  |
| [206. Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | [go/iterative](problems/algorithms/0206-reverse-linked-list/solutions/go/iterative/) | Reverse links in place | Save the next node, reverse the current link, then advance. | O(n) | O(1) |  |
| [206. Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | [go/recursive](problems/algorithms/0206-reverse-linked-list/solutions/go/recursive/) | Reverse the remaining list first | Reverse the suffix, point the next node back to the current node, then clear the current link. This is equivalent to using a stack iteratively. | O(n) | O(n), call stack |  |
| [238. Product of Array Except Self](problems/algorithms/0238-product-of-array-except-self/problem.md) | [go/prefix-suffix](problems/algorithms/0238-product-of-array-except-self/solutions/go/prefix-suffix/) | Exclude one index from a product; division is disallowed | Store exclusive prefix products in the result, then multiply by a running suffix product while scanning backward. | O(n) | O(1) auxiliary; output excluded |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/max-heap](problems/algorithms/0347-top-k-frequent-elements/solutions/go/max-heap/) | Select the k most frequent distinct values | Count frequencies, heapify value/count pairs, then extract the maximum k times. | O(n + k log m) expected | O(m) |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/min-heap](problems/algorithms/0347-top-k-frequent-elements/solutions/go/min-heap/) | Top k; retain only k candidates | Count frequencies, keep a size-k min-heap, and replace its root when a higher count appears. | O(n + m log k) expected | O(m) total; O(k) heap |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/bucket-sort](problems/algorithms/0347-top-k-frequent-elements/solutions/go/bucket-sort/) | Frequencies are bounded by the input length | Count frequencies and track the maximum; group values by count, then scan buckets downward until k values are collected. | O(n) expected | O(m + F), worst-case O(n) |  |
| [560. Subarray Sum Equals K](problems/algorithms/0560-subarray-sum-equals-k/problem.md) | [go/prefix-sum-hashmap](problems/algorithms/0560-subarray-sum-equals-k/solutions/go/prefix-sum-hashmap/) | Count contiguous ranges with an exact sum; values may be negative | For each prefix sum S, count earlier prefixes equal to S - k using a frequency map. | O(n) expected | O(n) | Started with O(n²); understood this approach with ChatGPT help. Updating the current prefix frequency before counting matches counts an empty range when k = 0: nums=[1] incorrectly returns 1 instead of 0. |
| [739. Daily Temperatures](problems/algorithms/0739-daily-temperatures/problem.md) | [go/monotonic-stack](problems/algorithms/0739-daily-temperatures/solutions/go/monotonic-stack/) | Find the next strictly greater value to the right | Keep unresolved indices in a stack with non-increasing temperatures. A warmer day pops indices and records their waiting times. | O(n) | O(n) |  |

<!-- generated:progress:end -->

## Using this repository

Each problem has a `problem.md` containing its statement, facts, and solution notes; a compact `testcases.json`; and its implementation files.

- [Conventions](docs/conventions.md): adding problems and updating the generated overview.
- [Testing and benchmarks](docs/testing-and-benchmarks.md): fixture format and measurement policy.

Edit `problem.md` to record your approach, trigger, key idea, complexity, and mistakes. The progress, index, and solution notes above are generated from those files.

Run `python3 scripts/consolidate.py` to refresh the overview, or let the pre-commit hook refresh it from staged documentation. Use `python3 scripts/consolidate.py --check` to check that it is current.

The pre-commit hook tests a temporary copy of all staged files before updating the README. Solution cases, shared runner tests, and repository automation tests must pass for the commit to proceed.

Go solutions share the root `go.mod`. Run `go run ./cmd/leetcode test` to generate adapters and check every Go implementation against its `testcases.json`, or see the [individual commands](docs/testing-and-benchmarks.md#running-tests). Benchmarking and the CI workflow are not implemented yet. Solution entries do not imply local verification or LeetCode acceptance.

Solution source files are maintained by the repository owner. Automated tooling must not edit, format, move, or rename them without explicit permission.

## License

My original solutions, tooling, and notes are licensed under the [MIT License](LICENSE). Third-party problem material is excluded and remains subject to its respective rights.
