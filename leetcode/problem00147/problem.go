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
func insertionSortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	// 最前面搞一个头节点，方便插入
	head = &ListNode{
		// Val:  math.MinInt64,
		Next: head,
	}

	// sortedNode 已经有序的最有一个节点，初始一个节点已经有序
	// sortedNode.Next 为下一个需要处理的节点
	for sortedNode := head.Next; nil != sortedNode.Next; {
		curNode := sortedNode.Next
		// 如果已经有序节点的最后一个，直接往后走，快速方法
		if curNode.Val >= sortedNode.Val {
			sortedNode = sortedNode.Next
			continue
		}
		// perNode->node->xxxx
		// 找到第一个节点，满足 node.Val >= insertNode.Val
		// 则插入到 node 前面，也就是 perNode 后面
		perNode := head
		for node := perNode.Next; node.Val < curNode.Val; perNode, node = node, node.Next {
		}

		// 取出 curNode 节点
		sortedNode.Next = curNode.Next
		// curNode 插入到 perNode 后面
		curNode.Next = perNode.Next
		perNode.Next = curNode
		// sortedNode 这个还是原来那个，不需要变化
	}

	return head.Next
}
