package problem00406

import (
	"sort"
)

func allCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	numPoints := R * C
	ans := make([][]int, numPoints)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ans[i*C+j] = []int{i, j}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		x, y := ans[i], ans[j]
		lx, ly := abs(x[0]-r0)+abs(x[1]-c0), abs(y[0]-r0)+abs(y[1]-c0)
		return lx < ly
	})
	return ans
}

func allCellsDistOrder2(R int, C int, r0 int, c0 int) [][]int {
	numPoints := R * C
	maxDistance := max(R-1-r0, r0) + max(C-1-c0, c0)
	// [0,maxDistance]
	buckets := make([][][]int, maxDistance+1)

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			l := abs(i-r0) + abs(j-c0)
			buckets[l] = append(buckets[l], []int{i, j})
		}
	}

	ans := make([][]int, 0, numPoints)
	for _, b := range buckets {
		ans = append(ans, b...)
	}
	return ans
}

func allCellsDistOrder3(R int, C int, r0 int, c0 int) [][]int {
	numPoints := R * C

	// * 首先网上移动一个，从上面开始
	// * 再向右下方向走，直到到达 r0 这一行
	// * 再向右下方向走，直到到达 c0 这一列
	// * 再向左上方向走，直到到达 r0 这一行
	// * 再向右上方向走，直到到达 c0 这一列 （此时回到原定，重新开始）
	dirs := [][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	ans := make([][]int, numPoints)
	ans[0] = []int{r0, c0}
	index := 1
	row, col := r0, c0
	for index < numPoints {
		row--
		for i, d := range dirs {
			for i&1 == 0 && row != r0 || i&1 == 1 && col != c0 {
				if 0 <= row && row < R && 0 <= col && col < C {
					ans[index] = []int{row, col}
					index++
				}
				row += d[0]
				col += d[1]
			}
		}
	}
	return ans
}

func allCellsDistOrder4(R int, C int, r0 int, c0 int) [][]int {
	numPoints := R * C

	// 初始放入 (r0,c0)
	queue := [][]int{{r0, c0}}
	// 记录，是否已经加入过
	visited := make([]bool, numPoints)
	visited[r0*C+c0] = true // 记录 (r0, c0) 已经入队了

	ans := make([][]int, numPoints)
	index := 0 // 出队的时候再统计

	for len(queue) > 0 {
		// 出队列
		point := queue[0]
		queue = queue[1:]

		// 记录结果
		ans[index] = point
		index++

		row, col := point[0], point[1]
		if row+1 < R && !visited[(row+1)*C+col] { // 上
			queue = append(queue, []int{row + 1, col})
			visited[(row+1)*C+col] = true
		}
		if row-1 >= 0 && !visited[(row-1)*C+col] { // 下
			queue = append(queue, []int{row - 1, col})
			visited[(row-1)*C+col] = true
		}
		if col-1 >= 0 && !visited[row*C+col-1] { // 左
			queue = append(queue, []int{row, col - 1})
			visited[row*C+col-1] = true
		}
		if col+1 < C && !visited[row*C+col+1] { // 右
			queue = append(queue, []int{row, col + 1})
			visited[row*C+col+1] = true
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
