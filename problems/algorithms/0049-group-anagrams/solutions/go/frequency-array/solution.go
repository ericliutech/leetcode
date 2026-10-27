package solution

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		var fmap [26]int
		for _, r := range s {
			fmap[r-'a']++
		}

		groups[fmap] = append(groups[fmap], s)
	}

	results := make([][]string, 0, len(groups))

	for _, v := range groups {
		results = append(results, v)
	}

	return results
}
