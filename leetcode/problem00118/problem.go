package problem00118

func generate(numRows int) [][]int {
	switch numRows {
	case 0:
		return nil
	case 1:
		return [][]int{{1}}
	case 2:
		return [][]int{{1}, {1, 1}}
	}
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
