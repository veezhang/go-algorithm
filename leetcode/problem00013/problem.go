package problem00013

func romanToInt(s string) int {
	romanTable := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	length := len(s)
	num := 0
	preNum := 0
	for i := 0; i < length; i++ {
		curNum := romanTable[s[i]]
		if preNum < curNum {
			num -= preNum
		} else {
			num += preNum
		}
		preNum = curNum
	}
	num += preNum

	return num
}
