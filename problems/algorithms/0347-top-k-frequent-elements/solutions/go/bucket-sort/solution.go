package solution

func topKFrequent(nums []int, k int) []int {
	fmap := make(map[int]int)

	maxFreq := 0
	for _, n := range nums {
		if fmap[n]+1 > maxFreq {
			maxFreq = fmap[n] + 1
		}
		fmap[n] = fmap[n] + 1
	}

	buckets := make([][]int, maxFreq+1)

	for num, freq := range fmap {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0 && len(result) < k; i-- {
		for j := 0; j < len(buckets[i]) && len(result) < k; j++ {
			result = append(result, buckets[i][j])
		}
	}

	return result
}
