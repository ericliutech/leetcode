package solution

func topKFrequent(nums []int, k int) []int {
    fmap := make(map[int]int)

    for _, n := range nums {
        fmap[n]++
    }

    maxheap := make([][2]int, 0, len(fmap))

    for n, freq := range fmap {
        maxheap = append(maxheap, [2]int{n, freq})
    }

    for i := len(maxheap)/2-1; i >= 0; i-- {
        siftDown(maxheap, i)
    }

    result := make([]int, k)
    for i := range k {
        result[i] = maxheap[0][0]
        maxheap[0] = maxheap[len(maxheap)-1]
        maxheap = maxheap[:len(maxheap)-1]
        siftDown(maxheap, 0)
    }

    return result
}

func siftDown(maxheap [][2]int, root int) {
    for {
        left := root*2+1

        if left >= len(maxheap) {
            break
        }

        right := left+1
        highest := root

        if maxheap[left][1] > maxheap[highest][1] {
            highest = left
        }
        if right < len(maxheap) && maxheap[right][1] > maxheap[highest][1] {
            highest = right
        }

        if highest == root {
            break
        }

        maxheap[highest], maxheap[root] = maxheap[root], maxheap[highest]

        root = highest
    }
}
