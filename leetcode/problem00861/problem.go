package problem00861

func matrixScore(A [][]int) (ans int) {
	// m行 * n列 的矩阵
	m, n := len(A), len(A[0])
	// 翻转行，让每一行最左侧为 1 ，则每一行此 1 对结果的共享为： 2^(n - 1)
	// 所以 m 行，则第一列的贡献为： 2^(n-1) * m
	ans = 1 << (n - 1) * m

	for col := 1; col < n; col++ {
		ones := 0
		// 遍历每一行，统计 1 的个数
		for i := 0; i < m; i++ {
			// 每行的第一个为 1 了， 所以后面与 A[i][0] 相等的也就是 1
			if A[i][col] == A[i][0] {
				ones++
			}
		}
		// 如果 1 比 0 少，则翻转列
		if ones < m-ones {
			ones = m - ones
		}
		// 第 col 行此 1 对结果的共享为： 2^(n - col - 1)
		// 所以 ones 个 1 ，贡献为： 2^(n - col - 1) * m
		ans += 1 << (n - col - 1) * ones
	}
	return
}
