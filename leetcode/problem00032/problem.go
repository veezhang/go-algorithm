package problem00032

func longestValidParentheses1(s string) (ans int) {
	if len(s) <= 1 {
		return
	}

	// dp[i] 表示以下标 i 字符结尾的最长有效括号的长度。以 '(' 结尾的子串对应的 dp 值必定为 0 。
	dp := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		// 只有 dp[i] == ')' 的时候，才会有值
		if s[i] != ')' {
			continue
		}

		if s[i-1] == '(' {
			// 1. 如果 dp[i-1] = '(' ，也就是类似 "..................()"，则 dp[i] = dp[i-1] + 2
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			// 2. 如果 s[i-1] = ')' ，也就是类似 "..................))" 这种的，则：
			//         只有当 dp[i-1] 为有效的括号对( dp[i-1] > 0 )的时候，dp[i] 才可能为有效的，类似 "......(...........))"
			//         只有当 s[i-1] = ')' 对应的位置(i - 1 - dp[i-1] + 1)必定为 '('
			//         并且只有当对应位置前面一个( i - 1 - dp[i-1] + 1 - 1 )也为 '(' ，s[i] 才为有效的括号对
			//         即： dp[i-1] > 0 && s[i - dp[i-1] - 1] == '(' 时候
			//         又由于 i - 1 - dp[i-1] + 1 - 1 之前可能也有有效的括号对，所以还需要加上，最终为：
			//         dp[i] = dp[i - dp[i - 1] - 2] + dp[i - 1] + 2
			if i-dp[i-1] >= 3 { // 前面至少要有两个字符，才可能有有效的括号对， i-dp[i-1]-2 >= 1
				dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
			} else {
				dp[i] = dp[i-1] + 2
			}
		}

		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return
}

func longestValidParentheses2(s string) (ans int) {
	if len(s) <= 1 {
		return
	}

	// 始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」
	// 初始化入栈 -1
	stack := []int{-1}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 1. 如果是 '(' ，入栈
			stack = append(stack, i)
		} else {
			// 2. 否则，出栈
			stack = stack[:len(stack)-1]
			if 0 == len(stack) {
				// 2.1 如果栈空了，说明当前的右括号为没有被匹配的右括号，则入栈
				stack = append(stack, i)
			} else {
				// 2.2 果栈不为空，则计算
				l := i - stack[len(stack)-1]
				if ans < l {
					ans = l
				}
			}
		}
	}
	return
}

func longestValidParentheses3(s string) (ans int) {
	if len(s) <= 1 {
		return
	}

	left, right, revLeft, revRight, length := 0, 0, 0, 0, len(s)
	for i := 0; i < length; i++ {
		// 正着统计下
		if '(' == s[i] {
			left++
		} else {
			right++
		}

		if left == right {
			l := left << 1 // left * 2
			if ans < l {
				ans = l
			}
		} else if right > left {
			left, right = 0, 0
		}

		// 反着统计下
		if '(' == s[length-1-i] {
			revLeft++
		} else {
			revRight++
		}
		if revLeft == revRight {
			l := revLeft << 1 // revLeft * 2
			if ans < l {
				ans = l
			}
		} else if revLeft > revRight {
			revLeft, revRight = 0, 0
		}
	}

	return
}
