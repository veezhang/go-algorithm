package problem01122

import (
	"math"
	"sort"
)

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	if 0 == len(arr1) {
		return arr1
	}

	if 0 == len(arr2) {
		sort.Ints(arr1)
		return arr1
	}

	posMap := make(map[int]int)
	for i, v := range arr2 {
		posMap[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		posX, xok := posMap[x]
		posY, yok := posMap[y]
		if xok && yok {
			return posX < posY
		}

		if xok || yok { // xok -> true ; !xok -> false
			return xok
		}

		return x < y
	})
	return arr1
}

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	if 0 == len(arr1) {
		return arr1
	}

	// 统计 arr1 中最小值和最大值
	min, max := math.MaxInt64, math.MinInt64
	for _, v := range arr1 {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// 统计 arr1 中出现的次数
	counts := make([]int, max-min+1)
	for _, v := range arr1 {
		counts[v-min]++
	}

	ans := make([]int, 0, len(arr1))
	// 遍历 arr2 ，先加入已经存在的数，加入后需要计数 -1
	for _, v := range arr2 {
		for ; counts[v-min] > 0; counts[v-min]-- {
			ans = append(ans, v)
		}
	}

	// 遍历 counts ，加入剩余的数字
	for v, n := range counts {
		for ; n > 0; n-- {
			ans = append(ans, v+min)
		}
	}

	return ans
}
