package problem00725

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	// 如果就只是分一个，直接返回
	if 1 == k {
		return []*ListNode{root}
	}

	// 统计链表长度
	length := 0
	for n := root; nil != n; n, length = n.Next, length+1 {
	}
	ans := make([]*ListNode, k)

	// sublen 表示子链的长度， remain 表示还有多少个剩余，没有分下去
	// 就可以算出来每个子链的长度
	sublen, remain := length/k, length%k
	// 如果只有一个链表，不需要拆分的情况
	if length <= sublen || remain > 0 && length <= sublen+1 {
		ans[0] = root
		return ans
	}

	// 每一个断点位置为：
	for i := 0; i < k; i++ {
		ans[i] = root
		if i == k-1 {
			// 最后不再需要了拆分了
			break
		}
		// 当前链的长度
		l := sublen
		if remain > 0 { // 如果还有剩余的，则需要 +1
			l++
			remain--
		}
		if l == 0 {
			break
		}
		for j := 0; j < l-1; j++ {
			root = root.Next
		}
		root, root.Next = root.Next, nil
	}
	return ans
}
