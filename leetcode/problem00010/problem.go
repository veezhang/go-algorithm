package problem00010

// 回溯法
func isMatchBackTrace(s string, p string) bool {
	// 如是 s 搜素完了，但是 p 还没有搜索完，则可能匹配（比如：s = "ab", p = "abc*d*e*f*"）
	// 也可能不匹配（比如： s = "ab", p = "abc"）
	// 所以这里结束判断依据 p 是否匹配完成

	// p 已经匹配完了后， 如果 s 也匹配完了，则是成功匹配；否则失败
	if len(p) == 0 {
		return len(s) == 0
	}

	// 第一个字符是否匹配， s[0],p[0]相等或者 p[0] 等于 '.'
	// 注意，此时 s 可能为空串了
	isFirstCharMatch := (len(s) > 0 && (s[0] == p[0] || p[0] == '.'))

	// 如果第二个字符是 '*' ， 也就是 x* 这种形式（ x 上面已经比较过）
	// 有两种情况， 第一 x* 匹配 0 次， 则直接 p 后移 2 位
	// 第二， x* 至少匹配 1 次数，且 isFirstCharMatch 必须等于 true，此时 s 后移一位，p 还可以继续匹配 s 中的
	if len(p) >= 2 && p[1] == '*' {
		// 如果 x* 表示 0 次失败后，回溯回来，看匹配 1 次是否成功
		return isMatchBackTrace(s, p[2:]) || (isFirstCharMatch && isMatchBackTrace(s[1:], p))
	} else {
		//此处不包含 * ，如果 isFirstCharMatch = true 则直接都后移一位
		return isFirstCharMatch && isMatchBackTrace(s[1:], p[1:])
	}
}

// 动态规划
// 时间复杂度： O(SP)
// 空间复杂度： O(SP)
func isMatchDP(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)

	dp := make([][]bool, lenS+1)

	for i := lenS; i >= 0; i-- {
		// 构造二维 slice
		dp[i] = make([]bool, lenP+1)

		// 初始情况
		// p 匹配完了，s 也匹配完了，，则为true ， 也就是 dp[lenS][lenP] = true ；
		// p 匹配完了，s 还没有匹配完了，则为 false ，， i 属于 [0, lenS -1] 时候，dp[i][lenP] = false ，不需要设置
		if i == lenS {
			dp[lenS][lenP] = true
		}

		// 依赖关系：
		//   ┌───────────────┬───────────────┐
		//   │ dp[i][j]      │ dp[i+1][j])   │
		//   ├───────────────┼───────────────┤
		//   │               │ dp[i+1][j+1]  │
		//   ├───────────────┼───────────────┤
		//   │ dp[i][j+2]    │               │
		//   └───────────────┴───────────────┘
		//

		for j := lenP - 1; j >= 0; j-- {
			// 第一个字符是否匹配， s[0],p[0]相等或者 p[0] 等于 '.'
			// 注意，此时 i 可能越界了
			isFirstCharMatch := (i < lenS && (s[i] == p[j] || p[j] == '.'))
			// 如果第二个字符是 '*' ， 也就是 x* 这种形式（ x 上面已经比较过）
			// 有两种情况， 第一 x* 匹配 0 次， 则直接 p 后移 2 位
			// 第二， x* 至少匹配 1 次数，且 isFirstCharMatch 必须等于 true，此时 s 后移一位，p 还可以继续匹配 s 中的
			if j+1 < lenP && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (isFirstCharMatch && dp[i+1][j])
			} else {
				//此处不包含 * ，如果 isFirstCharMatch = true 则直接都后移一位
				dp[i][j] = isFirstCharMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

// 动态规划 空间优化
// 时间复杂度： O(SP)
// 空间复杂度： O(P)
func isMatchDP1(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)

	// 依赖关系：
	//   ┌───────────────┬───────────────┐
	//   │ dp[i][j]      │ dp[i+1][j])   │
	//   ├───────────────┼───────────────┤
	//   │               │ dp[i+1][j+1]  │
	//   ├───────────────┼───────────────┤
	//   │ dp[i][j+2]    │               │
	//   └───────────────┴───────────────┘
	//
	// 如果空间优化为 dp := make([]bool, lenP+1) 表示第 i 列数据
	// 那么需要保证：
	// 1. 同一列的 d[j+2] 要先选出来
	// 2. 后一列的 d[j] 不能先计算了
	// 3. 后一列的 d[j+1] 不能先计算了
	// 满足 1 需要 j 从大到小来算，而满足 3 却要从小到大来算，所以矛盾，没有办法了
	//
	// 那么再申请一个 dp1 := make([]bool, lenP+1) 保留后一列的数据就可以了

	// dp 为当前要计算的列，也就是 第 i 列； dp1 为当前依赖的第 i+1 列
	dp := make([]bool, lenP+1)
	dp1 := make([]bool, lenP+1)

	for i := lenS; i >= 0; i-- {
		// 初始情况
		// p 匹配完了，s 也匹配完了，，则为true ， 也就是 dp[lenS][lenP] = true ；
		// p 匹配完了，s 还没有匹配完了，则为 false ，， i 属于 [0, lenS -1] 时候，dp[i][lenP] = false ，不需要设置
		// 由于这里重复利用了，所以 false 也需要设置，除了 1 == lenP 外，其他都是false
		if i == lenS {
			dp1[lenP] = true
		} else if dp1[lenP] {
			dp1[lenP] = false
		}
		// 交换下
		dp, dp1 = dp1, dp

		for j := lenP - 1; j >= 0; j-- {
			// 第一个字符是否匹配， s[0],p[0]相等或者 p[0] 等于 '.'
			// 注意，此时 i 可能越界了
			isFirstCharMatch := (i < lenS && (s[i] == p[j] || p[j] == '.'))
			// 如果第二个字符是 '*' ， 也就是 x* 这种形式（ x 上面已经比较过）
			// 有两种情况， 第一 x* 匹配 0 次， 则直接 p 后移 2 位
			// 第二， x* 至少匹配 1 次数，且 isFirstCharMatch 必须等于 true，此时 s 后移一位，p 还可以继续匹配 s 中的
			if j+1 < lenP && p[j+1] == '*' {
				// dp[i][j] = dp[i][j+2] || (isFirstCharMatch && dp[i+1][j])
				dp[j] = dp[j+2] || (isFirstCharMatch && dp1[j])
			} else {
				//此处不包含 * ，如果 isFirstCharMatch = true 则直接都后移一位
				// dp[i][j] = isFirstCharMatch && dp[i+1][j+1]
				dp[j] = isFirstCharMatch && dp1[j+1]
			}
		}
	}

	return dp[0]
}
