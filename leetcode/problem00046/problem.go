package problem00046

func permute1(nums []int) [][]int {
	var backtrack func(nums, pathNums []int, visited []bool, result *[][]int)
	backtrack = func(nums, pathNums []int, visited []bool, result *[][]int) {
		// 满足条件，加入到结果队列
		if len(nums) == len(pathNums) {
			*result = append(*result, append([]int(nil), pathNums...))
			return
		}

		for i := 0; i < len(nums); i++ {
			// 已经访问过
			if visited[i] {
				continue
			}
			// 做选择
			visited[i] = true
			pathNums = append(pathNums, nums[i])
			// backtrack(路径，选择列表)
			backtrack(nums, pathNums, visited, result)
			// 撤销选择
			pathNums = pathNums[:len(pathNums)-1]
			visited[i] = false
		}
	}
	pathNums := make([]int, 0, len(nums))
	visited := make([]bool, len(nums))
	result := make([][]int, 0)

	backtrack(nums, pathNums, visited, &result)

	return result
}

func permute2(nums []int) [][]int {
	// nFill 表示已经填充的个数
	var backtrack func(pathNums []int, nFill, length int, result *[][]int)
	backtrack = func(pathNums []int, nFill, length int, result *[][]int) {
		// 满足条件，加入到结果队列
		if nFill == length {
			*result = append(*result, append([]int(nil), pathNums...))
			return
		}

		// for 选择 in 选择列表
		for i := nFill; i < length; i++ {
			// 做选择
			pathNums[nFill], pathNums[i] = pathNums[i], pathNums[nFill]
			// backtrack(路径，选择列表)
			backtrack(pathNums, nFill+1, length, result)
			// 撤销选择
			pathNums[nFill], pathNums[i] = pathNums[i], pathNums[nFill]
		}
	}
	result := make([][]int, 0)
	backtrack(nums, 0, len(nums), &result)

	return result
}
