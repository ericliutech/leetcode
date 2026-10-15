package solution

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	ans := make([][]int, 0)

	i := 0
	for i < len(nums)-2 && nums[i] <= 0 {
		j, k := i+1, len(nums)-1

		for j < k && nums[k] >= 0 {
			sum := nums[i] + nums[j] + nums[k]

			if sum > 0 {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})

				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}

		i++
		for i < len(nums)-2 && nums[i] == nums[i-1] {
			i++
		}
	}

	return ans
}
