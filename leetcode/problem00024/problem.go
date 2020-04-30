package problem00024

// 递归
// 时间复杂度： O(N)
// 空间复杂度： O(N)
func swapPairsRecursion(head *ListNode) (ret *ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	//
	//   head			next
	//   ┌──────┐     ┌──────┐     ┌──────┐     ┌──────┐
	//   │ Val  │     │ Val  │     │ Val  │     │ Val  │
	//   ├──────┤     ├──────┤     ├──────┤     ├──────┤
	//   │ Next │     │ Next │     │ Next │     │ Next │
	//   └──────┘     └──────┘     └──────┘     └──────┘
	//
	// 需要替换 A 和 B
	// 最后 per -> B -> A -> C，然后 per  = A
	//
	// next := head.Next
	// head.Next = swapPairsRecursion(next.Next)
	// next.Next = head

	// return next

	ret, head.Next, head.Next.Next = head.Next, swapPairsRecursion(head.Next.Next), head
	return
}

// 指针操作
// 时间复杂度： O(N)
// 空间复杂度： O(1)
func swapPairsPointer(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{Next: head}
	//
	//   per			 A			  B            C
	//   ┌──────┐     ┌──────┐     ┌──────┐     ┌──────┐
	//   │ Val  │     │ Val  │     │ Val  │     │ Val  │
	//   ├──────┤     ├──────┤     ├──────┤     ├──────┤
	//   │ Next │     │ Next │     │ Next │     │ Next │
	//   └──────┘     └──────┘     └──────┘     └──────┘
	//
	// 需要替换 A 和 B
	// 最后 per -> B -> A -> C，然后 per  = A
	//
	// 引用中间变量保存 per.Next
	// a := per.Next            // 中间变量 a 保存下 per.Next
	// per.Next = per.Next.Next // per 下一个为 B
	// a.Next = per.Next.Next   // A 下一个为 C
	// per.Next.Next = a        // B 下一个为 A
	// per = a                  // per 移动到 a
	//
	// 也可以利用 go 多个值的复制
	// per, per.Next, per.Next.Next, per.Next.Next.Next = per.Next, per.Next.Next, per.Next.Next.Next, per.Next
	//
	per := ret

	for per.Next != nil && per.Next.Next != nil {
		// a := per.Next            // 中间变量 a 保存下 per.Next
		// per.Next = per.Next.Next // per 下一个为 B
		// a.Next = per.Next.Next   // A 下一个为 C
		// per.Next.Next = a        // B 下一个为 A
		// per = a                  // per 移动到 a

		per, per.Next, per.Next.Next, per.Next.Next.Next = per.Next, per.Next.Next, per.Next.Next.Next, per.Next
	}

	return ret.Next
}
