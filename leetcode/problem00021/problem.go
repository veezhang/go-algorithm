package problem00021

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
