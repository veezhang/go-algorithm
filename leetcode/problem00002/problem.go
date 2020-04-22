package problem00002

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		//重用下carry，计算当前节点的数字
		carry += v1 + v2

		curr.Next = &ListNode{
			Val: carry % 10,
		}
		curr = curr.Next
		carry /= 10
	}

	if carry != 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return result.Next
}
