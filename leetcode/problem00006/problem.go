package problem00006

func convert(s string, numRows int) string {
	length := len(s)
	if numRows < 2 || length <= numRows {
		return s
	}

	result := make([]byte, length)
	step := numRows<<1 - 2 // 2 * numRows - 2
	index := 0             // result中的下标
	i := 0                 // 原始字符串的下标

	// 遍历行
	for row := 0; row < numRows; row++ {
		// 遍历每一行的数据
		for i = row; i < length; i += step {
			result[index] = s[i]
			index++
			// 如果不是第一行和最后一行，每一行会多加一个字符
			// middle := i + step - row<<1 // (2*(numRows-row) - 2)
			if row != 0 && row != numRows-1 && i+step-row<<1 < length {
				result[index] = s[i+step-row<<1]
				index++
			}
		}
	}
	return string(result)
}
