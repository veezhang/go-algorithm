package problem00003

func lengthOfLongestSubstring(s string) int {
	result := 0
	length := len(s)

	// 记录各个字符的 下标+1
	var indices [128]int

	// 滑动窗口
	// 如果 s[j] 在 s[i,j]中，其下标是j1, 则i滑动到 j1 + 1, 也就是indices中的值，计算值，j滑动j++
	// 如果 s[j] 不在 s[i,j]中，计算值，j滑动j++
	for i, j := 0, 0; j < length; j++ {
		// s[j]不仅仅要存在，其下标需要>=i，下标+1 就>i
		if indices[s[j]] > i {
			i = indices[s[j]]
		}
		if j-i+1 > result {
			result = j - i + 1
		}
		indices[s[j]] = j + 1
	}

	return result
}
