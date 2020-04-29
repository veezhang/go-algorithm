package problem00022

// brackets 为当前添加了的括号的字符数组
// open 为开括号的数目， close 为闭括号的数目，n 为总括号的对数
// result 为存放结果的，这里使用指针，因为 append 会修改它
func generateParenthesisBacktrackOnce(brackets []byte, open, close, n int, result *[]string) {

	if open == n && close == n {
		// 此处不能使用零拷贝 *(*string)(unsafe.Pointer(&brackets))
		// 因为后面还需要使用 brackets ，这里用 string(brackets) 复制一份出来
		*result = append(*result, string(brackets))
		return
	}

	// 闭括号比开括号多，不可能匹配，返回
	if close > open {
		return
	}

	// 如果开括号还不够，增加一个
	if open < n {
		// 做选择： 	brackets[open+close] = '(' , open = open + 1
		brackets[open+close] = '('
		generateParenthesisBacktrackOnce(brackets, open+1, close, n, result)
		// 撤销选择：	open 还是之前的 open ， brackets[open+close] 不需要撤销，后面会覆盖的
	}

	// 如果开括号还不够，增加一个
	if close < n {
		// 做选择： 	brackets[open+close] = ')' , close = close - 1
		brackets[open+close] = ')'
		generateParenthesisBacktrackOnce(brackets, open, close+1, n, result)
		// 撤销选择：	close 还是之前的 close ， brackets[open+close] 不需要撤销，后面会覆盖的
	}
}

// 回溯法
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	ret := make([]string, 0)
	brackets := make([]byte, n*2)

	generateParenthesisBacktrackOnce(brackets, 0, 0, n, &ret)

	return ret
}
