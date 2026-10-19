package solution

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// mid := (left+right)/2 --> can overflow
		mid := (right-left)/2 + left

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// case 1: left side is sorted
			if nums[left] <= target && target < nums[mid] {
				// go left
				right = mid - 1
			} else {
				// go right
				left = mid + 1
			}
		} else {
			// case 2: right side is sorted
			if nums[mid] < target && target <= nums[right] {
				// go right
				left = mid + 1
			} else {
				// go left
				right = mid - 1
			}
		}
	}

	return -1
}
