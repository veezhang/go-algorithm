package problem00767

import (
	"container/heap"
)

// 大根堆，根最大，下面的元素都不大于根
type heapByteCount struct {
	ch    byte
	count int
}
type heapByteCounts []heapByteCount

func (h heapByteCounts) Len() int {
	return len(h)
}

func (h heapByteCounts) Less(i, j int) bool {
	// Less 决定是大根堆还是小根堆，这里为大根堆
	return h[i].count > h[j].count
}

func (h heapByteCounts) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapByteCounts) Push(x interface{}) {
	*h = append(*h, x.(heapByteCount))
}

func (h *heapByteCounts) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func reorganizeString(S string) string {
	length := len(S)
	if length <= 1 {
		return S
	}

	maxCount := 0
	countMap := make(map[byte]int, 0)
	chNum := 0
	for _, ch := range []byte(S) {
		cnt, ok := countMap[ch]
		if !ok {
			chNum++
		}
		cnt++
		countMap[ch] = cnt
		if maxCount < cnt {
			maxCount = cnt
			if maxCount > (length+1)/2 {
				// 某种字符就占据了一半以上的相同字符，则不满足
				// aaabb -> ababa	maxCount <= (5+1)/2 = 3 基数
				// aaab  -> aba		maxCount <= (4+1)/2 = 2 偶数
				return ""
			}
		}
	}

	h := make(heapByteCounts, 0, chNum)
	for ch, cnt := range countMap {
		h = append(h, heapByteCount{
			ch:    ch,
			count: cnt,
		})
	}
	heap.Init(&h)

	ans := make([]byte, 0, length)
	for h.Len() >= 2 {
		a, b := heap.Pop(&h).(heapByteCount), heap.Pop(&h).(heapByteCount)

		ans = append(ans, a.ch, b.ch)
		if a.count > 1 {
			a.count--
			heap.Push(&h, a)
		}
		if b.count > 1 {
			b.count--
			heap.Push(&h, b)
		}
	}
	if h.Len() > 0 {
		ans = append(ans, h.Pop().(heapByteCount).ch)
	}
	return string(ans)
}
