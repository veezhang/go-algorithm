package problem00118

func getRow(rowIndex int) []int {
	// rowIndex 是行索引
	switch rowIndex {
	case 0:
		return []int{1}
	case 1:
		return []int{1, 1}
	}
	// 索引为 rowIndex ，对应的数个数为： rowIndex+1
	// 规律： 后一行为前一行错位相加，例如
	//		1			1	1			1	2	1			1	3	3	1
	// +		1			1	1			1	2	1			1	3	3	1
	//		1	1		1	2	1		1	3	3	1		1	4	6	4	1
	ans := make([]int, rowIndex+1)
	ans[0], ans[1] = 1, 1
	for i := 2; i <= rowIndex; i++ {
		// 需要倒叙， j = 0 也不用管，一直为 1
		for j := i; j > 0; j-- {
			ans[j] += ans[j-1]
		}
	}

	return ans
}
