package problem00842

import "math"

func splitIntoFibonacci(S string) (F []int) {
	length := len(S)

	// index 为当前统计到的下标
	// sum 为前两个数的和
	// prev 为上一个数
	var bt func(index, sum, prev int) bool
	bt = func(index, sum, prev int) bool {
		// 如果处理到最后一个了，则看是否有三个数
		if index == length {
			return len(F) >= 3
		}

		cur := 0
		for i := index; i < length; i++ {
			if i != index && '0' == S[index] {
				// 每个块的数字一定不要以零开头，除非这个块是数字 0 本身
				break
			}
			// 数字统计
			cur = cur*10 + int(S[i]-'0')
			// 每个整数都符合 32 位有符号整数类型
			if cur > math.MaxInt32 {
				break
			}

			// F[i] + F[i+1] = F[i+2]
			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				// 如果大于，已经不满足了
				if cur > sum {
					break
				}
				// cur == sum
			}
			// 做选择
			F = append(F, cur)
			if bt(i+1, prev+cur, cur) {
				return true
			}
			// 撤销选择
			F = F[:len(F)-1]
		}

		// 这里不满足
		return false
	}
	bt(0, 0, 0)
	return
}
