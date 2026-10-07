package solution

func twoSum(nums []int, target int) []int {
	knownIndices := make(map[int]int, 0)
	for i, num := range nums {
		if i2, ok := knownIndices[target-num]; ok {
			return []int{i, i2}
		}

		knownIndices[num] = i
	}

	return []int{0, 1}
}
