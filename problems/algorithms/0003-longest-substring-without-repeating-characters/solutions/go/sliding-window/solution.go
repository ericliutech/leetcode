package solution

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := [256]int{}
	for i := range seen {
		seen[i] = -1
	}

	left := 0
	maxlen := 0

	for right := 0; right < len(s) && len(s)-left >= maxlen; right++ {
		pos := seen[s[right]]
		if pos >= left {
			left = pos + 1
		}
		seen[s[right]] = right

		if right-left+1 > maxlen {
			maxlen = right - left + 1
		}
	}

	return maxlen
}
