package problem00973

import (
	"container/heap"
	"sort"
)

func pointDistanceSquare(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func kClosest1(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	sort.Slice(points, func(i, j int) bool {
		pi, pj := points[i], points[j]
		return pointDistanceSquare(pi) < pointDistanceSquare(pj)
	})
	return points[:K]
}

// 大根堆，根最大，下面的元素都不大于根
type heapPoint struct {
	point          []int
	distanceSquare int // 缓存距离，避免重复计算
}
type heapPoints []heapPoint

func (h heapPoints) Len() int {
	return len(h)
}

func (h heapPoints) Less(i, j int) bool {
	// Less 决定是大根堆还是小根堆，这里为大根堆
	return h[i].distanceSquare > h[j].distanceSquare
}

func (h heapPoints) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapPoints) Push(x interface{}) {
	*h = append(*h, x.(heapPoint))
}

func (h *heapPoints) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest2(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	h := make(heapPoints, K)
	for i, point := range points[:K] {
		h[i] = heapPoint{point, pointDistanceSquare(point)}
	}
	heap.Init(&h)

	for _, point := range points[K:] {
		// 如果小于根，则替换根
		if dis := pointDistanceSquare(point); dis < h[0].distanceSquare {
			h[0] = heapPoint{point, pointDistanceSquare(point)}
			// h[0].distanceSquare = dis
			// h[0].point[0] = point[0]
			// h[0].point[1] = point[1]
			heap.Fix(&h, 0)
		}
	}

	// 获取前 K 个点
	result := make([][]int, 0, K)
	for _, hp := range h {
		result = append(result, hp.point)
	}
	return result
}

// 快排-填坑法
func kClosest3(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	// 选择基准值，并将其复制到最右边
	// 这里使用的是三数取中间值，顺便也调整了这三个数
	movePivotToHight := func(points [][]int, low, hight int) {
		// 剩余一个元素，或者空，直接返回
		if low <= hight+1 {
			return
		}
		mid := (low + hight) >> 1
		lowDis, midDis, hightDis := pointDistanceSquare(points[low]), pointDistanceSquare(points[mid]), pointDistanceSquare(points[hight])
		// 如果 lowDis > midDis 换下， 换完后则 points[low] <= points[mid]
		if lowDis > midDis {
			points[low], points[mid] = points[mid], points[low]
			lowDis, midDis = midDis, lowDis
		}
		// 如果 lowDis > hightDis 换下， 换完后则 points[low] <= points[hight]
		if lowDis > hightDis {
			points[low], points[hight] = points[hight], points[low]
			lowDis, hightDis = hightDis, lowDis
		}
		// 此时 points[low] 已经是最小的了， points[mid] 与 points[hight] 直接的顺序还未知
		// 而我们要做的是让 points[hight] 存的是中间值，而不是最大值，所以 midDis < hightDis 才交换
		if midDis < hightDis {
			points[mid], points[hight] = points[hight], points[mid]
			// midDis, hightDis = hightDis, midDis // 不用交换这个了
		}
	}

	// 分区
	partition := func(points [][]int, low, hight int) int {
		// 选取枢纽（基准值）， 然后与 hight 换一下
		movePivotToHight(points, low, hight)
		// 选取基准值，此时 points[hight] 空出来了
		pivot := points[hight]
		pivotDistanceSquare := pointDistanceSquare(pivot)

		for low < hight {
			// 从左边开始找，low 向后移动，找到第一个比基准值大的
			for low < hight && pointDistanceSquare(points[low]) < pivotDistanceSquare {
				low++
			}
			if low < hight {
				// 将第一个比基准值大的放到 points[hight] 中，此时 points[low] 空出来了
				points[hight] = points[low]
			}

			// 再先从右边开始找，hight 向前移动，找到第一个比基准值小的
			for low < hight && pointDistanceSquare(points[hight]) >= pivotDistanceSquare {
				hight--
			}
			if low < hight {
				// 第一个比基准值小的放到 points[low] 中，points[low] 为上一步空出来的
				points[low] = points[hight]
			}
		}
		// 此时 low == hight
		// 将基准值填入到 low 中
		points[low] = pivot
		return low
	}

	var quickSelect func(points [][]int, low, hight, K int)
	quickSelect = func(points [][]int, low, hight, K int) {
		if low >= hight {
			return
		}
		mid := partition(points, low, hight)
		// quickSelect(points, low, mid-1, K)
		// quickSelect(points, mid+1, hight, K)

		// 这里不是要全部排序号，只要满足左边 K 个值为最小的即可
		// mid 左边全部小于 mid ，右边全部大于，则 [0-mid] 全部是小于等于 points[mid] 的
		if mid+1 == K {
			return
		}

		if mid+1 > K {
			quickSelect(points, low, mid-1, K)
		} else {
			quickSelect(points, mid+1, hight, K)
		}
	}

	quickSelect(points, 0, len(points)-1, K)
	return points[:K]
}

// 快排-左右指针法
func kClosest4(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	// 选择基准值，并将其复制到最右边
	// 这里使用的是三数取中间值，顺便也调整了这三个数
	movePivotToHight := func(points [][]int, low, hight int) {
		// 剩余一个元素，或者空，直接返回
		if low <= hight+1 {
			return
		}
		mid := (low + hight) >> 1
		lowDis, midDis, hightDis := pointDistanceSquare(points[low]), pointDistanceSquare(points[mid]), pointDistanceSquare(points[hight])
		// 如果 lowDis > midDis 换下， 换完后则 points[low] <= points[mid]
		if lowDis > midDis {
			points[low], points[mid] = points[mid], points[low]
			lowDis, midDis = midDis, lowDis
		}
		// 如果 lowDis > hightDis 换下， 换完后则 points[low] <= points[hight]
		if lowDis > hightDis {
			points[low], points[hight] = points[hight], points[low]
			lowDis, hightDis = hightDis, lowDis
		}
		// 此时 points[low] 已经是最小的了， points[mid] 与 points[hight] 直接的顺序还未知
		// 而我们要做的是让 points[hight] 存的是中间值，而不是最大值，所以 midDis < hightDis 才交换
		if midDis < hightDis {
			points[mid], points[hight] = points[hight], points[mid]
			// midDis, hightDis = hightDis, midDis // 不用交换这个了
		}
	}

	// 分区
	partition := func(points [][]int, low, hight int) int {
		// 选取枢纽（基准值）， 然后与 hight 换一下
		movePivotToHight(points, low, hight)
		// 选取基准值，此时 points[hight] 空出来了
		pivot := points[hight]
		pivotDistanceSquare := pointDistanceSquare(pivot)

		// 记录原始的 hight
		ohight := hight

		for low < hight {
			// 从左边开始找，low 向后移动，找到第一个比基准值大的
			for low < hight && pointDistanceSquare(points[low]) <= pivotDistanceSquare {
				low++
			}
			// 再先从右边开始找，hight 向前移动，找到第一个比基准值小的
			for low < hight && pointDistanceSquare(points[hight]) >= pivotDistanceSquare {
				hight--
			}
			// 将比基准值大的和比基准值小的交换
			if low < hight {
				points[low], points[hight] = points[hight], points[low]
			}
		}
		// 此时 low == hight
		// 将其与基准值 points[ohight] 交换
		points[ohight], points[low] = points[low], points[ohight]
		return low
	}

	var quickSelect func(points [][]int, low, hight, K int)
	quickSelect = func(points [][]int, low, hight, K int) {
		if low >= hight {
			return
		}
		mid := partition(points, low, hight)
		// quickSelect(points, low, mid-1, K)
		// quickSelect(points, mid+1, hight, K)

		// 这里不是要全部排序号，只要满足左边 K 个值为最小的即可
		// mid 左边全部小于 mid ，右边全部大于，则 [0-mid] 全部是小于等于 points[mid] 的
		if mid+1 == K {
			return
		}

		if mid+1 > K {
			quickSelect(points, low, mid-1, K)
		} else {
			quickSelect(points, mid+1, hight, K)
		}
	}

	quickSelect(points, 0, len(points)-1, K)
	return points[:K]
}

// 快排-前后指针法
func kClosest5(points [][]int, K int) [][]int {
	if K >= len(points) {
		return points
	}

	// 选择基准值，并将其复制到最右边
	// 这里使用的是三数取中间值，顺便也调整了这三个数
	movePivotToHight := func(points [][]int, low, hight int) {
		// 剩余一个元素，或者空，直接返回
		if low <= hight+1 {
			return
		}
		mid := (low + hight) >> 1
		lowDis, midDis, hightDis := pointDistanceSquare(points[low]), pointDistanceSquare(points[mid]), pointDistanceSquare(points[hight])
		// 如果 lowDis > midDis 换下， 换完后则 points[low] <= points[mid]
		if lowDis > midDis {
			points[low], points[mid] = points[mid], points[low]
			lowDis, midDis = midDis, lowDis
		}
		// 如果 lowDis > hightDis 换下， 换完后则 points[low] <= points[hight]
		if lowDis > hightDis {
			points[low], points[hight] = points[hight], points[low]
			lowDis, hightDis = hightDis, lowDis
		}
		// 此时 points[low] 已经是最小的了， points[mid] 与 points[hight] 直接的顺序还未知
		// 而我们要做的是让 points[hight] 存的是中间值，而不是最大值，所以 midDis < hightDis 才交换
		if midDis < hightDis {
			points[mid], points[hight] = points[hight], points[mid]
			// midDis, hightDis = hightDis, midDis // 不用交换这个了
		}
	}

	// 分区
	partition := func(points [][]int, low, hight int) int {
		// 选取枢纽（基准值）， 然后与 hight 换一下
		movePivotToHight(points, low, hight)
		pivot := points[hight]
		pivotDistanceSquare := pointDistanceSquare(pivot)
		// storeIndex 记录下一个比基准值小的下标
		storeIndex := low

		for i := low; i < hight; i++ {
			// 如果比基准值小，则交换位置，并且 storeIndex++
			if pointDistanceSquare(points[i]) <= pivotDistanceSquare {
				points[i], points[storeIndex] = points[storeIndex], points[i]
				storeIndex++
			}
		}
		// 此时 [low, storeIndex) 全部小于基准值 points[hight]
		// 交换下 points[hight], points[storeIndex]
		points[hight], points[storeIndex] = points[storeIndex], points[hight]
		return storeIndex
	}

	var quickSelect func(points [][]int, low, hight, K int)
	quickSelect = func(points [][]int, low, hight, K int) {
		if low >= hight {
			return
		}
		mid := partition(points, low, hight)
		// quickSelect(points, low, mid-1, K)
		// quickSelect(points, mid+1, hight, K)

		// 这里不是要全部排序号，只要满足左边 K 个值为最小的即可
		// mid 左边全部小于 mid ，右边全部大于，则 [0-mid] 全部是小于等于 points[mid] 的
		if mid+1 == K {
			return
		}

		if mid+1 > K {
			quickSelect(points, low, mid-1, K)
		} else {
			quickSelect(points, mid+1, hight, K)
		}
	}

	quickSelect(points, 0, len(points)-1, K)
	return points[:K]
}
