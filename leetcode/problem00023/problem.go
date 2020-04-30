package problem00023

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}

	curr := ret

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return ret.Next
}

// K 个链表中找最小
// 时间复杂度： O(NK)， K 个链表，总共 N 个节点
// 空间复杂度： O(1)
func mergeKListsFindMinInK(lists []*ListNode) *ListNode {
	length := len(lists)

	switch length {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	ret := &ListNode{}
	curr := ret
	for {
		// min 表示最小值的链表
		var min int = -1
		for i := 0; i < length; i++ {
			if lists[i] == nil {
				continue
			}

			if min == -1 || lists[i].Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			break
		}
		curr.Next = lists[min]
		curr = curr.Next
		lists[min] = lists[min].Next
	}

	return ret.Next
}

// 一个一个的合并
// 时间复杂度： O(NK)， K 个链表，总共 N 个节点
// 空间复杂度： O(1)
func mergeKListsOneByOne(lists []*ListNode) *ListNode {
	var ret *ListNode
	for i := 0; i < len(lists); i++ {
		ret = mergeTwoLists(ret, lists[i])
	}
	return ret
}

// 分治法
// 时间复杂度： O(N*logK)
// 空间复杂度：	O(logK) 递归会使用到 O(logK) 空间代价的栈空间
func mergeKListsDivideAndConquer(lists []*ListNode) *ListNode {
	length := len(lists)

	switch length {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	}

	// 分别合并前一半和后一半，再合并到一起
	left := mergeKListsDivideAndConquer(lists[:length/2])
	right := mergeKListsDivideAndConquer(lists[length/2:])
	return mergeTwoLists(left, right)
}

// 分治法 非递归
// 时间复杂度： O(N*logK)
// 空间复杂度：	O(1)
func mergeKListsDivideAndConquerNoRecursion(lists []*ListNode) *ListNode {
	length := len(lists)

	switch length {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	}

	// k 为需要处理的链表个数，每次合并两个，如果偶数个，则剩余 k/2 ，如果是奇数个，最后有一个没有合并，则剩余 k/2 + 1
	// 也就是 (k + 1) / 2
	for k := length; k > 1; k = (k + 1) / 2 {
		for i := 0; i < k; i += 2 {
			// 最后只剩 1 个链表
			if i == k-1 {
				lists[i/2] = lists[i]
			} else {
				lists[i/2] = mergeTwoLists(lists[i], lists[i+1])
			}
		}
	}

	return lists[0]
}

// 小根堆
// 时间复杂度： O(N*logK)
// 空间复杂度： O(K)
func mergeKListsHeap(lists []*ListNode) *ListNode {
	length := len(lists)

	switch length {
	case 0:
		return nil
	case 1:
		return lists[0]
	}

	// 堆中最多有 length 个元素，这里直接分配，避免后面扩充
	h := make(Heap, 0, length)
	heap.Init(&h)

	for _, l := range lists {
		if l != nil {
			heap.Push(&h, l)
		}
	}

	ret := &ListNode{}
	curr := ret
	for h.Len() > 0 {
		// 获取堆最上面的元素，也就是当前堆中最小的元素
		// 由于链表是有序的，所以当前堆中最小的，也就是整个中最小的
		node := heap.Pop(&h).(*ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
	}
	return ret.Next
}

// 小根堆
// go/src/container/heap/heap.go
// type Interface interface {
// 	sort.Interface
// 	Push(x interface{}) // add x as element Len()
// 	Pop() interface{}   // remove and return element Len() - 1.
// }

type Heap []*ListNode

// 实现 sort.Interface
func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现 heap.Interface
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
