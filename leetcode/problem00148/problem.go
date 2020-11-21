package problem00147

import (
	. "github/veezhang/go-algorithm/leetcode/common/list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	// 统计链表长度
	length := 0
	for node := head; nil != node; node = node.Next {
		length++
	}

	// 最前面搞一个头节点，方便插入
	head = &ListNode{
		Next: head,
	}

	// 归并排序
	for subLen := 1; subLen < length; subLen <<= 1 {
		// preNode 已经处理了的节点的最后一个
		// curNode 当前处理的节点
		preNode, curNode := head, head.Next
		for nil != curNode {
			// 拆分第一个 subLen 长度的链表
			head1 := curNode
			// head1 为第一个，后面还需要 subLen - 1 个
			for i := 1; i < subLen && nil != curNode.Next; i, curNode = i+1, curNode.Next {
			}
			// 拆分第二个 subLen 长度的链表
			head2 := curNode.Next
			curNode.Next = nil // 断开链表
			curNode = head2
			if nil != curNode {
				// head2 为第一个，后面还需要 subLen - 1 个
				for i := 1; i < subLen && nil != curNode.Next; i, curNode = i+1, curNode.Next {
				}

				// 如果后面还有需要处理的，更新 curNode
				if nil != curNode {
					// curNode.Next = nil 断开链表
					curNode, curNode.Next = curNode.Next, nil
				}
			}

			// 合并两个 subLen 长度的链表
			// 并且更新 preNode
			preNode.Next, preNode = mergeTwoLists(head1, head2)
		}
	}

	return head.Next
}

// problem00021 类似
// 返回第一个和最后一个 *ListNode
func mergeTwoLists(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode) {
	ret := &ListNode{}

	curr := ret
	last := ret

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		if nil == curr.Next {
			last = curr
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	for nil != last.Next {
		last = last.Next
	}

	return ret.Next, last
}
