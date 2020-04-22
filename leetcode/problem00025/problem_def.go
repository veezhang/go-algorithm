package problem00025

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewFromSlice(s []int) *ListNode {

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

func (l *ListNode) ToSlice() []int {
	result := []int{}

	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}

	return result
}

func (l *ListNode) Equal(t *ListNode) bool {

	for {
		if (l == nil) != (t == nil) {
			return false
		}

		if l == nil {
			return true
		}

		if l.Val != t.Val {
			return false
		}

		l = l.Next
		t = t.Next
	}
}
