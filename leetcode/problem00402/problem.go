package problem00402

func removeKdigits(num string, k int) string {
	// 如果全部移除
	if k >= len(num) {
		return "0"
	}

	// 栈
	stack := []byte{}

	// 遍历每一个字符
	for i := 0; i < len(num); i++ {
		v := num[i]
		// 如果还需要继续移除元素，并且栈顶比当前的大，则出栈
		for k > 0 && len(stack) > 0 && v < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		if 0 == k { // 如果已经全部移除完了，把后面的加上，break
			stack = append(stack, num[i:]...)
			break
		}
		// 当前入栈
		stack = append(stack, v)
	}

	// 如果还有没有移除的，则后面的清除，因为 stack 是递增的
	if k > 0 {
		stack = stack[:len(stack)-k]
	}
	// 清除前面的 '0'
	k = 0
	for ; len(stack) > k && '0' == stack[k]; k++ {
	}
	if k > 0 {
		stack = stack[k:]
	}

	// 如果为全部删除了
	if 0 == len(stack) {
		return "0"
	}

	return string(stack)
}
