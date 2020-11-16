package problem00406

import (
	"sort"
)

func reconstructQueue1(people [][]int) [][]int {
	if len(people) <= 1 {
		return people
	}

	// 将所有人按照身高从低到高排序，然后一个一个的安排其位置。
	// 如果前面有 k 个人高于或者等于这个人，则需要空出 k 个位置出来，插入到第 k + 1 个空位置。
	// 为了统一算法，一样高的人，因为相同高度的也会算 k 值 ，
	// 也就是后面出现相同升高的人，可以认为比前面的矮一丢丢，所有 k 按照降序排序。
	// 举例说明：
	// 先排好序
	// [[4,4], [5,2], [5,0], [6,1], [7,1], [7,0]]
	// 先插入 [4,4] ， 前面需要空出 4 个位置
	// [     ,      ,      ,      , [4,4],      ]
	// 再插入 [5,2] ， 前面需要空出 0 个位置
	// [     ,      , [5,2],      , [4,4],      ]
	// 再插入 [5,0] ， 前面需要空出 2 个位置
	// [[5,0],      , [5,2],      , [4,4],      ]
	// 再插入 [6,1] ， 前面需要空出 1 个位置
	// [[5,0],      , [5,2], [6,1], [4,4],      ]
	// 再插入 [7,1] ， 前面需要空出 1 个位置
	// [[5,0],      , [5,2], [6,1], [4,4], [7,1]]
	// 再插入 [7,0] ， 前面需要空出 0 个位置
	// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

	sort.Slice(people, func(i, j int) bool {
		x, y := people[i], people[j]
		return x[0] < y[0] || x[0] == y[0] && x[1] > y[1]
	})

	ans := make([][]int, len(people))
	minNilIndex := 0 // 记录最小的可能为空位置，前面一定全部填充了
	for _, p := range people {
		k := p[1] + 1
		b := false // 表示 minNilIndex 是否已经更新了
		for i := minNilIndex; i < len(ans); i++ {
			if nil == ans[i] {
				if !b {
					minNilIndex = i
					b = true
				}
				k--
				if 0 == k {
					ans[i] = p
					if minNilIndex == i {
						minNilIndex++
					}
					break
				}
			}
		}
	}
	return ans
}

func reconstructQueue2(people [][]int) [][]int {
	if len(people) <= 1 {
		return people
	}

	// 将所有人按照身高从高到低降序排序，然后一个一个的安排其位置。
	// 由于只有比其高的人才会影响安排的位置，而比其高的又全部已经先站好位置了
	// 所有只要插入到 k 个人后面即可，也就是插入到第 k + 1 个位置，也就是插入到 pepole[i] 。
	// 为了统一算法，一样高的人，因为相同高度的也会算 k 值 ，
	// 也就是后面出现相同升高的人，可以认为比前面的矮一丢丢，所有 k 按照升序排序。
	// 举例说明：
	// 先排好序
	// [[7,0], [7,1], [6,1], [5,0], [5,2], [4,4]]
	// 先插入 [7,0] ，插入到 [0]
	// [[7,0],      ,      ,      ,     ,      ]
	// 先插入 [7,1] ， 插入到 [1]
	// [[7,0], [7,1],      ,      ,     ,      ]
	// 先插入 [6,1] ， 插入到 [1]
	// [[7,0], [6,1], [7,1],      ,      ,     ]
	// 先插入 [5,0] ， 插入到 [0]
	// [[5,0], [7,0], [6,1], [7,1],      ,     ]
	// 先插入 [5,2] ， 插入到 [2]
	// [[5,0], [7,0], [6,1], [7,1],      ,     ]
	// 先插入 [5,2] ， 插入到 [2]
	// [[5,0], [7,0], [5,2], [6,1], [7,1],     ]
	// 先插入 [4,4] ， 插入到 [4]
	// [[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]

	sort.Slice(people, func(i, j int) bool {
		x, y := people[i], people[j]
		return x[0] > y[0] || x[0] == y[0] && x[1] < y[1]
	})
	ans := make([][]int, len(people))
	lastIdx := -1 // 已经插入的最后一个元素
	for _, p := range people {
		k := p[1]

		// 将 ans[k:] 向后移动一步
		if lastIdx >= k {
			// [k+1,lastIdx+1] => [k+1,lastIdx+2)
			copy(ans[k+1:lastIdx+2], ans[k:])
		}
		// 插入到 ans[k]
		ans[k] = p
		// 更新下 lastIdx
		if lastIdx < k {
			lastIdx = k
		} else {
			lastIdx++
		}
	}

	return ans
}
