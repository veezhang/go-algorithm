package problem00136

func singleNumber(nums []int) int {
	// nums 是非空，就不做判断了
	s := 0
	for _, n := range nums {
		s ^= n
	}
	return s
}
