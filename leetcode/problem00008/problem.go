package problem00008

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

// 普通解法
func myAtoiNormal(str string) int {
	length := len(str)
	i := 0

	// 过滤空格
	for ; i < length && str[i] == ' '; i++ {
	}

	// 全是空格或者空串
	if i == length {
		return 0
	}

	num := 0
	sign := 1
	if str[i] == '-' || str[i] == '+' {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}

	for ; i < length && str[i] >= '0' && str[i] <= '9'; i++ {
		d := int(str[i] - '0')
		if sign == 1 {
			if num > MAX_INT32_DIV_10 || (num == MAX_INT32_DIV_10 && d > MAX_INT32_MOD_10) {
				return math.MaxInt32
			}
		} else {
			if num > 0 {
				num = -num
			}
			d = -d
			if num < MIN_INT32_DIV_10 || (num == MIN_INT32_DIV_10 && d < MIN_INT32_MOD_10) {
				return math.MinInt32
			}
		}
		num = num*10 + d
	}

	return num
}

// 有限状态机 deterministic finite automaton, DFA
func myAtoiDFA(str string) int {
	if "" == str {
		return 0
	}

	// |           |     ' '   |    +/-    |  number   |  other    |
	// |-----------|-----------|-----------|-----------|-----------|
	// | start     | start     | signed    | in_number | end       |
	// | signed    | end       | end       | in_number | end       |
	// | in_number | end       | end       | in_number | end       |
	// | end       | end       | end       | end       | end       |

	type status_type int

	const (
		STATUS_START status_type = iota
		STATUS_SIGNED
		STATUS_IN_NUMBER
		STATUS_END
	)

	var statusTable = map[status_type][]status_type{
		STATUS_START:     {STATUS_START, STATUS_SIGNED, STATUS_IN_NUMBER, STATUS_END},
		STATUS_SIGNED:    {STATUS_END, STATUS_END, STATUS_IN_NUMBER, STATUS_END},
		STATUS_IN_NUMBER: {STATUS_END, STATUS_END, STATUS_IN_NUMBER, STATUS_END},
		STATUS_END:       {STATUS_END, STATUS_END, STATUS_END, STATUS_END},
	}

	getCharCol := func(c byte) int {
		switch c {
		case ' ':
			return 0
		case '-', '+':
			return 1
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return 2
		default:
		}
		return 3
	}

	num := 0
	sign := 1
	status := STATUS_START

outer:
	for i := 0; i < len(str); i++ {
		status = statusTable[status][getCharCol(str[i])]
		switch status {
		case STATUS_START:
			// nothing to do
		case STATUS_SIGNED:
			if str[i] == '-' {
				sign = -1
			}
		case STATUS_IN_NUMBER:
			d := int(str[i] - '0')
			if sign == 1 {
				if num > MAX_INT32_DIV_10 || (num == MAX_INT32_DIV_10 && d > MAX_INT32_MOD_10) {
					num = math.MaxInt32
					break outer
				}
			} else {
				if num > 0 {
					num = -num
				}
				d = -d
				if num < MIN_INT32_DIV_10 || (num == MIN_INT32_DIV_10 && d < MIN_INT32_MOD_10) {
					num = math.MinInt32
					break outer
				}
			}
			num = num*10 + d
		case STATUS_END:
			break outer
		}
	}

	return num
}
