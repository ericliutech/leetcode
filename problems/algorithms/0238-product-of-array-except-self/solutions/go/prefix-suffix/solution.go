package solution

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	// populate product prefixes
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// multiply by suffix which is compounded on the go
	base := 1
	for i := len(nums) - 2; i >= 0; i-- {
		base *= nums[i+1]
		result[i] *= base
	}

	return result
}
