package problem00842

import (
	"strconv"
)

func monotoneIncreasingDigits1(N int) int {
	s := []byte(strconv.Itoa(N))
	length := len(s)
	last9 := length - 1
	for i := length - 1; i > 0; i-- {
		// 如果前面的数大一些，则前面的数字--，后面的都为 9
		if s[i-1] > s[i] {
			s[i-1]--
			for j := i; j <= last9; j++ {
				s[j] = '9'
			}
			last9 = i - 1
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}

func monotoneIncreasingDigits2(N int) (ans int) {
	lo, hi, base := 0, 0, 1
	for 0 != N {
		lo = N % 10
		N /= 10
		hi = N % 10

		// 如果低位小于高位，则后面的全部为 9
		// N 需要借一位
		if lo < hi {
			ans = base*10 - 1
			N-- // N 需要借一位
		} else {
			ans += lo * base
		}
		base *= 10
	}
	return
}

func monotoneIncreasingDigits3(N int) (ans int) {
	ones := 111111111
	// 数可以改为 a0 * 111111111 + a1 * 11111111 + ... + an * 1
	for i := 0; i < 9; i++ {
		for ans+ones > N {
			ones /= 10
		}
		ans += ones
		if ans == N {
			break
		}
	}
	return
}
