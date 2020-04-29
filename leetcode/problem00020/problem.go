package problem00020

func isValid(s string) bool {
	length := len(s)

	if length == 0 {
		return true
	}
	// 字符对应的 ASCII 码
	// ( 40
	// ) 41
	// [ 91
	// ] 93
	// { 123
	// } 125
	//
	// 经测试, slice 比 map 快很多，所以这里我们使用 slice ，会浪费部分内存，如果考虑内存可以使用
	// Benchmark_ByteMapSliceSwitch/map-4              59312358                19.3  ns/op
	// Benchmark_ByteMapSliceSwitch/slice-4            746922408                1.53 ns/op
	// Benchmark_ByteMapSliceSwitch/switch-4           465211090                2.51 ns/op

	// 闭括号对应的映射，这里减最小的字符 '(' ，可以缩减 closeBracketMap 大小
	closeBracketMap := []byte{
		')' - '(': '(',
		']' - '(': '[',
		'}' - '(': '{',
	}

	// 避免 append 扩充内存带来的性能影响，这里也直接长度为 length
	stack := make([]byte, length, length)
	curIndex := 0
	for i := 0; i < length; i++ {
		// 如果是闭括号
		// 获取 s[i] 对应的开括号，如果没有返回 0 ，则是 s[i] 自己都是开括号，否则 s[i] 是闭括号
		siOpenBracket := closeBracketMap[s[i]-'(']

		// 如果 s[i] 是开括号，入栈
		if 0 == siOpenBracket {
			stack[curIndex] = s[i]
			curIndex++
		} else {
			// 如果 s[i] 是闭括号

			// 但是 s[i] 对应的开括号不等于栈顶元素，直接返回 false
			if curIndex == 0 || siOpenBracket != stack[curIndex-1] {
				return false
			}
			// 否则，出栈。减小就好，不需要清理对象
			curIndex--
		}
	}

	// 栈是否为空
	return curIndex == 0
}
