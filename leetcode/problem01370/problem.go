package problem00222

func sortString(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}

	// 统计最小最大的字符串
	minChar, maxChar := int('z'), int('a')
	for _, ch := range []byte(s) {
		ch := int(ch)
		if ch > maxChar {
			maxChar = ch
		}
		if ch < minChar {
			minChar = ch
		}
	}

	// 全部是同一个字符
	if minChar == maxChar {
		return s
	}

	// 统计词频
	countSlice := make([]int, maxChar-minChar+1)
	for _, ch := range []byte(s) {
		ch := int(ch)
		countSlice[int(ch-minChar)]++
	}
	ans := make([]byte, length)
	index := 0
	for min, max := 0, maxChar-minChar; index < length; {
		// 从小到大获取
		for i := min; i <= max; i++ {
			if countSlice[i] > 0 {
				ans[index] = byte(i + minChar)
				countSlice[i]--
				index++
			}
			// 统计下最小的字符，更新，避免后面无效的查找
			if i == min && 0 == countSlice[i] {
				min++
			}
		}
		// 从大到小获取
		for i := max; i >= min; i-- {
			if countSlice[i] > 0 {
				ans[index] = byte(i + minChar)
				countSlice[i]--
				index++
			}
			// 统计下最大的字符，更新，避免后面无效的查找
			if i == max && 0 == countSlice[i] {
				max--
			}
		}
	}

	return string(ans)
}
