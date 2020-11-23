package problem00242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countMap := make(map[rune]int, len(s))
	for _, r := range s {
		countMap[r]++
	}

	for _, r := range t {
		countMap[r]--
		if countMap[r] < 0 {
			return false
		}

	}
	return true
}
