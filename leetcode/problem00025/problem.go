package problem00025

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if k == 1 {
		return head
	}

	result := &ListNode{}
	resultTail := result

	//K子链的head
	for head != nil {
		//找K个子链，head已经是一个了，再找k-1个
		headK := head
		tailK := head
		for i := 0; i < k-1; i++ {
			if tailK.Next == nil {
				//将最后不满足的添加到后面
				resultTail.Next = headK
				return result.Next
			}
			tailK = tailK.Next
		}

		head = tailK.Next

		//断开K个节点
		tailK.Next = nil

		//反转，完成后revK为之前K子链的最后一个，tailK现在变成反转后的最后一个
		revK := reverseList(headK)

		//将K子链连接到result
		resultTail.Next = revK
		resultTail = headK
	}

	return result.Next
}

func reverseList(head *ListNode) *ListNode {
	var result *ListNode

	for head != nil {
		next := head.Next
		head.Next = result
		result = head
		head = next
	}

	return result
}
