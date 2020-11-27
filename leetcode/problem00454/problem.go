package problem00164

func fourSumCount(A []int, B []int, C []int, D []int) (ans int) {
	if 0 == len(A) {
		return
	}

	if 1 == len(A) {
		if 0 == A[0]+B[0]+C[0]+D[0] {
			return 1
		}
		return
	}
	countMap := make(map[int]int, len(A))
	// 统计 a+b 的信息，可以得到相同结果的组合数目
	for _, a := range A {
		for _, b := range B {
			countMap[a+b]++
		}
	}

	// 再在 C,D 中找到 a + b + c + d = 0 的
	// -(c + d) = a + b
	for _, c := range C {
		for _, d := range D {
			ans += countMap[-(c + d)]
		}
	}
	return
}
