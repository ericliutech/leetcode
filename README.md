# LeetCode solutions

My LeetCode problems, solutions, and notes for future review.

<!-- generated:progress:start -->

## Progress

| Measure | Count |
| --- | ---: |
| Documented problems | 9 |
| Solution entries | 12 |

Counts reflect documentation, not correctness or LeetCode acceptance.

## Problem index

| ID | Problem | Difficulty | Topics | Languages |
| --- | --- | --- | --- | --- |
| 1 | [Two Sum](problems/algorithms/0001-two-sum/problem.md) | Easy | Array, Hash table | go |
| 3 | [Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | Medium | Hash table, String, Sliding window | go |
| 15 | [3Sum](problems/algorithms/0015-3sum/problem.md) | Medium | Array, Two pointers, Sorting | go |
| 20 | [Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | Easy | String, Stack | go |
| 33 | [Search in Rotated Sorted Array](problems/algorithms/0033-search-in-rotated-sorted-array/problem.md) | Medium | Array, Binary search | go |
| 56 | [Merge Intervals](problems/algorithms/0056-merge-intervals/problem.md) | Medium | Array, Sorting | go |
| 102 | [Binary Tree Level Order Traversal](problems/algorithms/0102-binary-tree-level-order-traversal/problem.md) | Medium | Tree, Breadth-first search, Binary tree | go |
| 206 | [Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | Easy | Linked list, Recursion | go |
| 347 | [Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | Medium | Array, Hash table, Heap, Bucket sort | go |

## Solution notes

| Problem | Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- | --- |
| [1. Two Sum](problems/algorithms/0001-two-sum/problem.md) | [go/hashmap](problems/algorithms/0001-two-sum/solutions/go/hashmap/) | Pair sum; fast complement lookup | Check `target-num` among earlier values, then store `num → index`. | O(n) expected | O(n) |  |
| [3. Longest Substring Without Repeating Characters](problems/algorithms/0003-longest-substring-without-repeating-characters/problem.md) | [go/sliding-window](problems/algorithms/0003-longest-substring-without-repeating-characters/solutions/go/sliding-window/) | Longest contiguous segment with no duplicates | Track last-seen indices; move `left` past a duplicate inside the window and update the maximum length. | O(n) | O(1), fixed 256-entry array | Tried recursive DP; a sliding window tracks uniqueness directly. |
| [15. 3Sum](problems/algorithms/0015-3sum/problem.md) | [go/two-pointer](problems/algorithms/0015-3sum/solutions/go/two-pointer/) | Zero-sum triplets; sorting guides pair search | Sort, fix one value, move two pointers by the sum, and skip duplicates. | O(n²) | O(log n) sorting stack | A hash-based pair search complicates duplicate handling; sorting lets me skip duplicates directly. |
| [20. Valid Parentheses](problems/algorithms/0020-valid-parentheses/problem.md) | [go/stack](problems/algorithms/0020-valid-parentheses/solutions/go/stack/) | Nested pairs; last opened must close first | Push opening brackets. For each closing bracket, require a matching stack top, then pop. Accept only if the stack ends empty. | Θ(n) worst case | O(n) |  |
| [33. Search in Rotated Sorted Array](problems/algorithms/0033-search-in-rotated-sorted-array/problem.md) | [go/binary-search](problems/algorithms/0033-search-in-rotated-sorted-array/solutions/go/binary-search/) | Rotated sorted array; O(log n) required | Identify the sorted half, then use its bounds to decide which half can contain the target. | O(log n) | O(1) | Tried recursion; iteration avoids extra stack space. |
| [56. Merge Intervals](problems/algorithms/0056-merge-intervals/problem.md) | [go/sort-and-merge](problems/algorithms/0056-merge-intervals/solutions/go/sort-and-merge/) | Overlapping ranges; sorted starts make merging local | Sort by start, merge into the last kept interval, and compact disjoint intervals into the same slice prefix. | O(n log n) | O(log n) sorting stack | Tried brute force; repeated slice splicing then made the merge pass O(n²). |
| [102. Binary Tree Level Order Traversal](problems/algorithms/0102-binary-tree-level-order-traversal/problem.md) | [go/bfs](problems/algorithms/0102-binary-tree-level-order-traversal/solutions/go/bfs/) | Visit a tree level by level | Process the current queue length, enqueue left then right children, and clear and drop the processed prefix. | O(n) | O(w) auxiliary queue |  |
| [206. Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | [go/iterative](problems/algorithms/0206-reverse-linked-list/solutions/go/iterative/) | Reverse links in place | Save the next node, reverse the current link, then advance. | O(n) | O(1) |  |
| [206. Reverse Linked List](problems/algorithms/0206-reverse-linked-list/problem.md) | [go/recursion](problems/algorithms/0206-reverse-linked-list/solutions/go/recursion/) | Reverse the remaining list first | Reverse the suffix, point the next node back to the current node, then clear the current link. This is equivalent to using a stack iteratively. | O(n) | O(n), call stack |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/max-heap](problems/algorithms/0347-top-k-frequent-elements/solutions/go/max-heap/) | Select the k most frequent distinct values | Count frequencies, heapify value/count pairs, then extract the maximum k times. | O(n + k log m) expected | O(m) |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/min-heap](problems/algorithms/0347-top-k-frequent-elements/solutions/go/min-heap/) | Top k; retain only k candidates | Count frequencies, keep a size-k min-heap, and replace its root when a higher count appears. | O(n + m log k) expected | O(m) total; O(k) heap |  |
| [347. Top K Frequent Elements](problems/algorithms/0347-top-k-frequent-elements/problem.md) | [go/bucket-sort](problems/algorithms/0347-top-k-frequent-elements/solutions/go/bucket-sort/) | Frequencies are bounded by the input length | Count frequencies and track the maximum; group values by count, then scan buckets downward until k values are collected. | O(n) expected | O(m + F), worst-case O(n) |  |

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
