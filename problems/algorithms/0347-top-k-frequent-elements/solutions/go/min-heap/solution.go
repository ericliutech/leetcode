package solution

func topKFrequent(nums []int, k int) []int {
	fmap := make(map[int]int)

	for _, n := range nums {
		fmap[n]++
	}

	minheap := make([][2]int, 0, k)

	for n, freq := range fmap {
		if len(minheap) < k {
			// min-heap not full; push
			minheap = append(minheap, [2]int{n, freq})

			siftUp(minheap, len(minheap)-1)
		} else if freq > minheap[0][1] {
			// min-heap full, incoming is larger, replace min
			minheap[0] = [2]int{n, freq}

			siftDown(minheap, 0)
		}
	}

	result := make([]int, k)
	for idx, tup := range minheap {
		result[idx] = tup[0]
	}

	return result
}

func siftUp(mh [][2]int, idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2

		if mh[idx][1] < mh[parent][1] {
			mh[idx], mh[parent] = mh[parent], mh[idx]

			idx = parent
		} else {
			break
		}
	}
}

func siftDown(mh [][2]int, idx int) {
	for {
		left := idx*2 + 1

		if left >= len(mh) {
			break
		}

		right := left + 1
		highest := idx

		if mh[left][1] < mh[highest][1] {
			highest = left
		}
		if right < len(mh) && mh[right][1] < mh[highest][1] {
			highest = right
		}

		if highest == idx {
			break
		}

		mh[highest], mh[idx] = mh[idx], mh[highest]

		idx = highest
	}
}
