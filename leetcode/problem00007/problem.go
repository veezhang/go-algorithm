package problem00007

import (
	"math"
)

const (
	// math.MaxInt32 1<<31 - 1		+2147483647
	// math.MinInt32 -1 << 31		-2147483648
	MAX_INT32_DIV_10 = math.MaxInt32 / 10 // 214748364
	MAX_INT32_MOD_10 = math.MaxInt32 % 10 // 7
	MIN_INT32_DIV_10 = math.MinInt32 / 10 // -214748364
	MIN_INT32_MOD_10 = math.MinInt32 % 10 // -8
)

func reverse(x int) int {
	rev := 0
	for x != 0 {
		mod := x % 10
		x /= 10
		// 溢出情况：
		// 	1.rev * 10 + mod > math.MaxInt32
		//   	1.1 rev > math.MaxInt32/10 则溢出
		//		1.2 rev == math.MaxInt32/10 && mod > math.MaxInt32 % 10 则溢出
		// 	2.rev * 10 + mod < math.MinInt32
		//   	1.1 rev < math.MinInt32/10 则溢出
		//		1.2 rev == math.MinInt32/10 && mod < math.MinInt32 % 10 则溢出
		if rev > MAX_INT32_DIV_10 || (rev == MAX_INT32_DIV_10 && mod > MAX_INT32_MOD_10) {
			return 0
		}

		if rev < MIN_INT32_DIV_10 || (rev == MIN_INT32_DIV_10 && mod < MIN_INT32_MOD_10) {
			return 0
		}

		rev = rev*10 + mod
	}

	return rev
}
