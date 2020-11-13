package problem00328

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) Slice() []int {
	s := []int{}
	for n := head; nil != n; n = n.Next {
		s = append(s, n.Val)
	}
	return s
}

func NewListFromSlice(s []int) *ListNode {
	head := &ListNode{}
	temp := head
	for i := 0; i < len(s); i++ {
		temp.Next = &ListNode{
			Val: s[i],
		}
		temp = temp.Next
	}

	return head.Next
}
