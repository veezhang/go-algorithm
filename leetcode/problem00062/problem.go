package problem00062

func uniquePaths1(m int, n int) int {
	var ans uint64 = 1
	for x, y := uint64(n), uint64(1); y < uint64(m); x, y = x+1, y+1 {
		ans = ans * x / y
	}
	return int(ans)
}

func uniquePaths2(m int, n int) int {
	if m > n {
		m, n = n, m
	}
	// 保证 m 为较小者， m*n 和 n*m 结果一致

	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}
	// 先遍历列
	for j := 1; j < n; j++ {
		// 在遍历行
		for i := 1; i < m; i++ {
			dp[i] += dp[i-1]
		}
	}

	return dp[m-1]
}
