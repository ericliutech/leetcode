package solution

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	// sort first to avoid interim gaps while merging
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// last is the last element we want
	// avoids splicing the slice when removing duplicate elements
	last := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[last][1] {
			// overlap found, update ending bound (merged)
			intervals[last][1] = max(intervals[last][1], intervals[i][1])
		} else {
			// advance last pointer but also copy the values to evaluate next
			last++
			intervals[last] = intervals[i]
		}
	}

	return intervals[:last+1] // only return what we want
}
