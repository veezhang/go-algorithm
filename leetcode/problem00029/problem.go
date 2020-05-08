package problem00029

import "math"

// 测试32位溢出情况，直接再定义一个函数，参数返回值都是 int32 类型
func divide32(dividend int32, divisor int32) (result int32) {
	switch divisor {
	case 1:
		return dividend
	case -1:
		if dividend == math.MinInt32 {
			// 溢出
			// −2^31 / -1 = 2^31
			return math.MaxInt32
		}
		return -dividend
	}
	// 判断是否同号
	// 符号位 1 为 负，0 为正，只有符号位不相同的时候，异或后才为 1 (负数)
	sameSign := dividend^divisor >= 0

	// 都转为负数求解，因为 −2^31 转为正数是会溢出的
	// 也可以转为正数，前提是需要多一些临界值判断
	if dividend > 0 {
		dividend = -dividend
	}

	if divisor > 0 {
		divisor = -divisor
	}

	// 此处开始， dividend 和 divisor 都为负数

	switch {
	case dividend > divisor: // |dividend| < |divisor| ，用绝对值可能更加符合人类比较，更容易看出来，下同
		return 0 // 为 0
	case dividend == divisor:
		if sameSign {
			return 1
		}
		return -1
	}

	// |dividend| >= |divisor|
	for dividend <= divisor {
		n := 0
		// dividend = 2^0 * divisor + 2^1 * divisor + ... + 2^n * divisor
		// 这里求最大的 n ，满足 |dividend| >= |2^n * divisor| 并且 |dividend| < |2^(n+1) * divisor|、
		// 为什么这里 divisor<<(n+1) 是 n+1 ? 因为满足条件后 n++ 了，所以就是 n + 1
		//
		// 因为 divisor<<(n+1) 可能越界，所以判断 divisor<<n >= math.MinInt32>>1 ， 那么 divisor<<(n+1) >= math.MinInt32 ，不会越界了
		//
		for divisor<<n >= math.MinInt32>>1 && dividend <= divisor<<(n+1) {
			n++
		}
		// result 加上 2^n
		result += 1 << n
		// dividend 减去 divisor << n
		dividend -= divisor << n
	}

	if sameSign {
		return result
	}

	return -result
}

func divide(dividend int, divisor int) (result int) {
	return int(divide32(int32(dividend), int32(divisor)))
}
