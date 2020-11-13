package problem00328

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList1(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	// odd 表示奇数节点， even 表示偶数节点。
	odd, even, evenHead := head, head.Next, head.Next
	// 循环拆分为奇偶节点
	for nil != even && nil != even.Next {
		odd.Next = even.Next // 基数节点下一个 = 偶数节点下一个为基数节点
		odd = odd.Next       // 基数节点向后移动
		even.Next = odd.Next // 偶数节点下一个 = 基数节点下一个为偶数节点
		even = even.Next     // 偶数节点向后移动
	}
	// 合并起来，基数后面为偶数
	odd.Next = evenHead

	return head
}

func oddEvenList2(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	// orderedOdd 表示已经有序的基数节点， nextOddPre 表示下一个基数节点的前一个。
	orderedOdd, nextOddPre := head, head.Next
	// 循环直到没有偶数节点了
	for nil != nextOddPre && nextOddPre.Next != nil {
		nextOdd := nextOddPre.Next
		// 1. 把下一个 nextOdd 取出来
		nextOddPre.Next = nextOdd.Next
		// 2. 把 nextOdd 插入到 orderedOdd 后面
		orderedOdd.Next, nextOdd.Next = nextOdd, orderedOdd.Next
		// 3. orderedOdd 向后移动一步，nextOddPre 向后移动两步（由于之前已经拿走了一个，所以只需要移动一步了）
		orderedOdd = nextOdd
		nextOddPre = nextOddPre.Next
	}

	return head
}
