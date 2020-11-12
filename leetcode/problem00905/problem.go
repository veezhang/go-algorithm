package problem00905

func sortArrayByParity1(A []int) []int {
	// i 用来遍历，j 为偶数的下标
	for i, j := 0, 0; i < len(A); i++ {
		// 如果 A[i] 为偶数
		if 0 == A[i]&1 {
			// 则交换
			if i != j {
				A[i], A[j] = A[j], A[i]
			}
			j++
		}
	}

	return A
}

func sortArrayByParity2(A []int) []int {
	// i 用来遍历，j 为偶数的下标
	for i, j := 0, len(A)-1; i < j; {
		// A[i] 为基数， A[j] 为偶数
		if A[i]&1 > A[j]&1 {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
		// A[i] 为偶数，继续后面找
		for ; i < j && 0 == A[i]&1; i++ {
		}
		// A[i] 为基数，继续后前找
		for ; i < j && 1 == A[j]&1; j-- {
		}
	}

	return A
}
