package problem00240

func searchMatrixStep(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])

	for r, c := 0, col-1; r < row && c >= 0; {
		if matrix[r][c] == target {
			return true
		}

		if matrix[r][c] > target {
			c--
		} else {
			r++
		}
	}
	return false
}

func searchMatrixSubMatrixProcess(matrix [][]int, target, left, up, right, down int) bool {
	if left > right || up > down {
		return false
	}

	if target < matrix[up][left] || target > matrix[down][right] {
		return false
	}

	mid := left + (right-left)>>1
	row := up
	for row <= down && matrix[row][mid] <= target {
		if matrix[row][mid] == target {
			return true
		}
		row++
	}

	return searchMatrixSubMatrixProcess(matrix, target, left, row, mid-1, down) || searchMatrixSubMatrixProcess(matrix, target, mid+1, up, right, down)
}

func searchMatrixSubMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	col := len(matrix[0])

	return searchMatrixSubMatrixProcess(matrix, target, 0, 0, col-1, row-1)
}
