package problem00007

// 最高位和个位比
func isPalindromeCompareHightAndLow(x int) bool {
	// x 为负数， 则不可能
	// x 不等于 0 并且个位为 0 也不可能
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// 获取有多少位
	// 比如 12321 -> 12321 / 10000 = 1 退出， pow10 = 10000
	pow10 := 1
	for x/pow10 >= 10 {
		pow10 *= 10
	}

	// 获取数字的最高位数字与个位相比，然后再去除这两个位，直到为 0 ，或不相等
	// 比如： 12321
	// 1. 	12321 / 10000 == 12321 % 10
	// 		12321 % 10000 / 10 = 232
	// 		10000 / 100 = 100
	// 2. 	232 / 100 == 232 % 10
	// 		232 % 100 / 10 = 3
	// 		100 / 100 = 1
	// 3.	3 / 1 = 3 % 10
	// 		3 % 1 / 10 = 0

	for x != 0 {
		if x/pow10 != x%10 {
			return false
		}
		x = x % pow10 / 10
		pow10 /= 100
	}

	return true
}

// 倒转一半数字
func isPalindromeReverseHalf(x int) bool {
	// x 为负数， 则不可能
	// x 不等于 0 并且个位为 0 也不可能
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// 比如： 123321， 倒转后面的 123 再与 前面的 123 相比较
	// 比如： 1234321，倒转后面的 1234/10 再与 前面的 123 相比较， 奇数 最后要过滤中间的 4
	// 因为只是倒装一半，所以不可能越界
	reverseHalfOfX := 0
	for x > reverseHalfOfX {
		reverseHalfOfX = reverseHalfOfX*10 + x%10
		x /= 10
	}

	return x == reverseHalfOfX || x == reverseHalfOfX/10
}
