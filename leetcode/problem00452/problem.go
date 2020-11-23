package problem00452

import (
	"sort"
)

func findMinArrowShots(points [][]int) (nArrow int) {
	length := len(points)
	if length <= 1 {
		return length
	}

	// 按右边界来排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// 记录当前气球最右边的位置，初始需要一只箭
	maxRight := points[0][1]
	nArrow = 1
	// 从后面的气球开始遍历
	for _, p := range points[1:] {
		// 如果气球的左边 > maxRight(之前气球的最右边)
		// 则重新记录 maxRight ，需要一只箭把之前的射掉
		if p[0] > maxRight {
			maxRight = p[1]
			nArrow++
		}
	}
	return
}
