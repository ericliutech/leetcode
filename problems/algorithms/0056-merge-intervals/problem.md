# 56. Merge Intervals

https://leetcode.com/problems/merge-intervals/

Difficulty: Medium
Topics: Array, Sorting
Comparison: unordered

The statement below is paraphrased.

## Statement

Each entry in `intervals` contains the inclusive start and end of a range. Combine overlapping ranges so the result covers exactly the same points, with no overlaps remaining. Ranges that share an endpoint must also be combined.

## Examples

| Case | `intervals` | Expected ranges |
| --- | --- | --- |
| 1 | `[[1, 3], [2, 6], [8, 10], [15, 18]]` | `[[1, 6], [8, 10], [15, 18]]` |
| 2 | `[[1, 4], [4, 5]]` | `[[1, 5]]` |
| 3 | `[[4, 7], [1, 4]]` | `[[1, 7]]` |

These examples and additional edge cases are stored in [testcases.json](testcases.json).

## Constraints

- Input size: 1 through 10,000 intervals, inclusive.
- Each interval contains exactly two integers.
- Bounds satisfy 0 ≤ start ≤ end ≤ 10,000.

## Output rules

Compare the collection of merged intervals without regard to outer order. Within each interval, the start and end retain their positions. This solution returns the intervals in increasing start order.

## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/sort-and-merge](solutions/go/sort-and-merge/) | Overlapping ranges; sorted starts make merging local | Sort by start, merge into the last kept interval, and compact disjoint intervals into the same slice prefix. | O(n log n) | O(log n) sorting stack | Tried brute force O(n³); repeated slice splicing then made the merge pass O(n²). |

`n` is the input interval count. Sorting takes O(n log n); the merge scan takes O(n) and O(1) auxiliary space. The current Go toolchain uses pattern-defeating quicksort with a heapsort fallback and O(log n) stack space for `slices.SortFunc`.

The earlier brute-force attempt was estimated as worse than O(n³). Even after sorting, repeatedly deleting or replacing entries through `append` could shift O(n) remaining slice headers per merge, producing O(n²) total work. Taking a subslice alone is O(1); copying the tail while rejoining is the costly part.

`last` is a write index, while `i` scans the input. Copying `intervals[i]` into `intervals[last]` copies an inner slice header, not all its values. Each input interval is examined once, and `intervals[:last+1]` returns the kept prefix without copying it.

The solution sorts and mutates its input; returned intervals share the input's storage. Empty input is outside the problem's constraints and is not handled by this implementation.
