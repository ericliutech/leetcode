package solution

func subarraySum(nums []int, k int) int {
	count := 0

	prefix := make(map[int]int, len(nums)+1)
	prefix[0] = 1

	rollingSum := 0

	for _, curr := range nums {
		rollingSum += curr // obtain rolling sum up to curr

		count += prefix[rollingSum-k] // count earlier prefix sums that would leave a subarray sum of k

		prefix[rollingSum] += 1 // update frequencies of previously encountered prefix sums
	}

	return count
}
