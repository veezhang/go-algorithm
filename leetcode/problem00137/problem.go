package problem00137

func singleNumber1(nums []int) int {
	// nums 是非空，就不做判断了
	bits1, bits2 := 0, 0
	for _, n := range nums {
		// 看 bits1 中那些已经置 1 的位，在 n 里面又出现了，则出现 2 次，记录下来
		bits2 |= (bits1 & n)
		// 将 n 记录到 bits1 里
		bits1 ^= n
		// bits1 和 bits2 里面都出现的位，说明出现了 3 次
		bits3 := bits1 & bits2
		// 出现 3 次，bits1 和 bits2 对应的位清零
		bits1 &= ^bits3
		bits2 &= ^bits3
	}

	return bits1
}

func singleNumber(nums []int) int {
	// nums 是非空，就不做判断了

	bits1, bits2 := 0, 0 // 出现一次的位，和两次的位

	for _, n := range nums {
		// 既不在出现一次的 bits1，也不在出现两次的 bits2 里面
		// 记录下来，出现了一次，再次出现则会抵消
		bits1 = (bits1 ^ n) & ^bits2
		// 既不在出现两次的 bits2 里面，也不在出现一次的 bits1 里面(已经计算了 bits1 ，还是不在，表示不止一次了)
		// 记录出现两次，第三次则会抵消
		bits2 = (bits2 ^ n) & ^bits1
	}
	return bits1
}
