package problem00260

func singleNumber(nums []int) []int {
	// nums 是非空，就不做判断了
	xXORy := 0
	for _, n := range nums {
		xXORy ^= n
	}

	diffBit := xXORy & (-xXORy)
	x := 0
	for _, n := range nums {
		if 0 == diffBit&n {
			x ^= n
		}
	}

	return []int{x, xXORy ^ x}
}
