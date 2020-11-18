package problem00134

func canCompleteCircuit(gas []int, cost []int) int {
	// curSum: 表示从 start 开始，到当前的剩余的油量
	// totalSum: 表示总油量
	// start: 表示起点
	curSum, totalSum, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		// 从 start 无法走到到 i ，则中间的任何点都无法走到 i
		// 用下一个点 i+1 开始，并清空 curSum
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}

	// 总油量不够
	if totalSum < 0 {
		return -1
	}

	return start
}
