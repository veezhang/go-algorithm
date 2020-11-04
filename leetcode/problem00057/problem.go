package problem00057

func insert(intervals [][]int, newInterval []int) (ans [][]int) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	merged := false        // newInterval 是否合并过
	minCalculated := false // 有交集的时候，是否已经计算过最小值了，第一次计算过后，后面不能比之前小了
	for _, interval := range intervals {
		if interval[0] > newInterval[1] {
			// 没有交集，interval 在 newInterval 右边
			// newInterval[0] newInterval[1] interval[0] interval[1]
			if !merged {
				// 都已经扫描到右边了，如果还没有合并，则合并下
				merged = true
				ans = append(ans, newInterval)
			}
			ans = append(ans, interval)
		} else if interval[1] < newInterval[0] {
			// 没有交集，interval 在 newInterval 左边
			// interval[0] interval[1] newInterval[0] newInterval[1]

			// 还没有到 newInterval ，直接加 interval 即可
			ans = append(ans, interval)
		} else {
			// 有交集，需要求并集
			// 注意后面可能还有交集，这里更新 newInterval
			if !minCalculated {
				newInterval[0] = min(interval[0], newInterval[0])
			}
			newInterval[1] = max(interval[1], newInterval[1])
		}
	}
	if !merged {
		ans = append(ans, newInterval)
	}
	return
}
