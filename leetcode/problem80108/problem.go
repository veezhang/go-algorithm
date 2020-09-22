package problem80108

func setZeroes(matrix [][]int) {
	col0 := false

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0 //下表为 r 的行包含 0
				if c == 0 {
					col0 = true // 下表为 0 的列包含 0
				} else {
					matrix[0][c] = 0 //下表为 c 的列包含 0
				}
			}
		}
	}

	// 由于下表为 0 的行和列存储了相关信息，所以，需要最后处理
	for r := len(matrix) - 1; r >= 0; r-- {
		for c := len(matrix[0]) - 1; c >= 0; c-- {
			if c > 0 {
				if matrix[r][0] == 0 || matrix[0][c] == 0 {
					matrix[r][c] = 0
				}
			} else if col0 {
				matrix[r][c] = 0
			}
		}
	}
}
