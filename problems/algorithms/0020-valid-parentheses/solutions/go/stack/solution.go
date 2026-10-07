package solution

func isValid(s string) bool {
	isClosing := map[rune]bool{
		'}': true,
		']': true,
		')': true,
	}

	matchingOpening := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	st := make([]rune, 0, len(s)/2)

	for _, r := range s {
		if isClosing[r] {
			if len(st) > 0 && st[len(st)-1] == matchingOpening[r] {
				st = st[:len(st)-1]
			} else {
				return false
			}
		} else {
			st = append(st, r)
		}
	}

	return len(st) == 0
}
