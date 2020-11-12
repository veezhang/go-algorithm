package problem00922

func sortArrayByParityII1(A []int) []int {
	// i 表示偶数下标， j 表示基数下标
	for i, j := 0, 1; i < len(A); i += 2 {
		// 如果偶数下标的数不是偶数，则需要迁移
		if 0 != A[i]&1 {
			// 找到第一个基数下标不是基数的位置
			for 0 != A[j]&1 { // 是基数，继续往后
				j += 2
			}
			// 然后交换
			A[i], A[j] = A[j], A[i]
		}
	}

	return A
}

func sortArrayByParityII2(A []int) []int {
	// i 表示偶数下标， j 表示基数下标
	for i, j := 0, len(A)-1; i < len(A) && j >= 0; {
		// A[i] 为基数， A[j] 为偶数
		if A[i]&1 > A[j]&1 {
			A[i], A[j] = A[j], A[i]
			i += 2
			j -= 2
		}
		// A[i] 为偶数，继续后面找
		for ; i < len(A) && j >= 0 && 0 == A[i]&1; i += 2 {
		}
		// A[i] 为基数，继续后前找
		for ; i < len(A) && j >= 0 && 1 == A[j]&1; j -= 2 {
		}
	}

	return A
}
