package solution

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	slices.Sort(candidates)

	var backtrack func(remain int, start int)
	backtrack = func(remain int, start int) {
		// valid path - add to result
		if remain == 0 {
			result = append(result, slices.Clone(path))
			return
		}

		// Your loop logic lives here:
		for i := start; i < len(candidates); i++ {
			// skip futile paths
			if candidates[i] > remain {
				break
			}

			// add [i] to path
			path = append(path, candidates[i])

			// backtrack
			backtrack(remain-candidates[i], i)

			// skip current candidate
			path = path[:len(path)-1]

		}
	}

	backtrack(target, 0)
	return result
}
