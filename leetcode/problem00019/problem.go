package problem00019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 先设置新的节点用来指向结果链表的头节点
	ret := &ListNode{Next: head}
	fast, slow := ret, ret

	// 快指针从列表的开头向前移动 n+1 步，然后快指针和慢指针一起移动，
	// 快指针和慢指针相隔 n 个节点。
	// 快指针到达 nil 的时候，慢指针还是保持相差 n 个节点
	// 比如： ...->n->n-1...->3->2->1->nil
	// 则慢指针的下一个就是倒是第 n 个节点

	for i := 0; i <= n; i++ {
		// 题意确保 n 是有效的，这里就不做长度不满足 n 的情况，也就不判断 fast 为 nil 的情况
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return ret.Next
}
