package problem00014

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	// 遍历第一个字符串，看后面的是否等于这个前缀
	for s0, len0, i0 := strs[0], len(strs[0]), 0; i0 < len0; i0++ {
		for i := 1; i < length; i++ {
			// 如果越界 strs[i] 或者 不想等了
			if i0 >= len(strs[i]) || s0[i0] != strs[i][i0] {
				return s0[:i0]
			}
		}
	}

	return strs[0]
}
