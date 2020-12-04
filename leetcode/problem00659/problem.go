package problem00164

import (
	"container/heap"
	"sort"
)

// int 小根堆
type intHeap struct{ sort.IntSlice }

func (h *intHeap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *intHeap) Pop() interface{} {
	old := h.IntSlice
	n := len(old)
	x := old[n-1]
	h.IntSlice = old[0 : n-1]
	return x
}

func (h *intHeap) push(v int) {
	heap.Push(h, v)
}

func (h *intHeap) pop() int {
	return heap.Pop(h).(int)
}

func isPossible1(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}

	// mHeapLens 中 key 表示： 当前数字为； value 表示： 以当前元素为最后一个数的长度组成的最小堆
	mHeapLens := make(map[int]*intHeap, 0)
	for _, n := range nums {
		var curLen int
		if h := mHeapLens[n-1]; h != nil {
			// 如果以前一个数为结尾的数存在
			// 则取出一个最小长度的和当前的拼接
			preLen := h.pop()
			if 0 == h.Len() {
				delete(mHeapLens, n-1)
			}
			// 和当前数拼接
			curLen = preLen + 1
		} else {
			// 此时已经中断了，检测下是否已经不满足了
			// 此时，只能存在为 n 的数，其它的都可以清理了
			// 另外，如果还有不是以 n 为最后一个元素的长度小于 3 的，则不满足条件，返回 false
			var hn *intHeap
			for i, h := range mHeapLens {
				if i == n {
					hn = h
				} else if h.IntSlice[0] < 3 {
					return false
				}
			}
			// 清理数据
			mHeapLens = make(map[int]*intHeap, 0)
			if nil != hn {
				mHeapLens[n] = hn
			}
			// 将当数据添加，长度为 1
			curLen = 1
		}

		// 如果当前数未统计过，则添加下
		if nil == mHeapLens[n] {
			mHeapLens[n] = &intHeap{}
		}
		mHeapLens[n].push(curLen)
	}

	for _, h := range mHeapLens {
		if h.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}

func isPossible2(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}

	// 每个数字的剩余个数
	leftCount := make(map[int]int, 0)
	for _, n := range nums {
		leftCount[n]++
	}
	// 以某个数字结尾的连续子序列的个数
	endCount := make(map[int]int, 0)

	for _, n := range nums {
		if 0 == leftCount[n] {
			continue
		}

		if endCount[n-1] > 0 { // 若存在以 n-1 结尾的连续子序列，则将 n 加到其末尾
			leftCount[n]--
			endCount[n-1]--
			endCount[n]++
		} else if leftCount[n+1] > 0 && leftCount[n+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
			leftCount[n]--
			leftCount[n+1]--
			leftCount[n+2]--
			endCount[n+2]++
		} else { // 否则不满足
			return false
		}
	}

	return true
}

func isPossible3(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}

	// len1Count, len2Count, len3Count 分别为子序列长度为 1,2 >=3 的个数
	len1Count, len2Count, len3Count := 0, 0, 0

	for i := 0; i < length; {
		// 统计连续相等的数字的个数
		start := i
		for i++; i < length && nums[i] == nums[start]; i++ {
		}
		// 统计完成，相同的数字区间为 [start, i)，个数为 i - start
		count := i - start

		if start == 0 { // 如果是第 1 个数及其相同的，则设置 len1Count 即可
			len1Count = count
		} else if nums[start-1]+1 < nums[start] { // 如果 start 前一个数 nums[start-1] 和当前数 nums[start] 不连续，中断了
			if len1Count > 0 || len2Count > 0 {
				// 此时，如果还有长度为 1 或者 2 的子序列，则不满足
				return false
			}
			// 此时重新组成新序列， 长度为 1 的为当前相等的数字的个数，长度为 2 的为0，长度为 3 的清零
			len1Count, len3Count = count, 0
		} else { // 此时是连续的
			// count 个数，优先拼接短的， 先 len1Count ，在 len2Count ，最后再 len3Count (加长原有的链)
			// 如果 len3Count 分不完，则多的重新组成新长度为 1 的子序列，也就是给 len1Count
			// 但是，必须有足够的来拼接 len1Count 和 len2Count
			leftCount := count - len1Count - len2Count
			if leftCount < 0 {
				// 此时 len1Count, len2Count 过多， count 都不足以把所有的都给加长，则不满足
				return false
			}

			if leftCount <= len3Count {
				// 剩余的不足以和所有的 len3Count 接起来，则
				// len3Count = len2Count+leftCount 	==> 之前的 len2Count 都连接了一个数，
				// 										然后 leftCount 个数把原来的 len3Count 也连接了一个数
				// 										其它 len3Count 中没有连接的断了也没事儿，哪些已经满足了
				// len2Count = len1Count 			==> 之前的 len1Count 都连接了一个数
				// len1Count = 0 					==> 没有多余的组成单个节点的子序列
				len1Count, len2Count, len3Count = 0, len1Count, len2Count+leftCount
			} else {
				// 剩余的不足以和所有的 len3Count 接起来，还有多的需要单组组成一个节点的子序列，则
				// len3Count = len3Count+leftCount 	==> 之前的 len2Count 都连接了一个数，
				// 										然后所有的 len3Count 也都连接了一个数
				// len2Count = len1Count 			==> 之前的 len1Count 都连接了一个数
				// len1Count = leftCount-len3Count 	==> 还有多的需要单组组成一个节点的子序列
				len1Count, len2Count, len3Count = leftCount-len3Count, len1Count, len2Count+len3Count
			}
		}
	}

	// 只要判断还有没有 1 个节点和 2 个节点的子序列即可
	return len1Count == 0 && len2Count == 0
}
