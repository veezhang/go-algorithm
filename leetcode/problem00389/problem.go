package problem00714

func findTheDifference1(s string, t string) byte {
	var diff int
	for i := 0; i < len(s); i++ {
		diff += int(t[i] - s[i])
	}

	diff += int(t[len(t)-1])

	return byte(diff)
}

func findTheDifference2(s string, t string) (diff byte) {
	for i := 0; i < len(s); i++ {
		diff ^= s[i] ^ t[i]
	}

	diff ^= t[len(t)-1]
	return
}
