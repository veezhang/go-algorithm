package problem00129

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbersDFS(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	// 进位加
	sum = sum*10 + root.Val
	if nil == root.Left && nil == root.Right {
		// 没有子树
		return sum
	}
	// 再遍历左子树和右子树
	return dfs(root.Left, sum) + dfs(root.Right, sum)
}

func sumNumbersBFS(root *TreeNode) int {
	if nil == root {
		return 0
	}

	type QNode struct {
		TreeNode *TreeNode
		Sum      int // 到此 TreeNode 为止，前面已经累计的数
	}

	queue := list.New()
	queue.PushBack(QNode{root, root.Val})

	totalSum := 0

	for queue.Len() > 0 {
		node := queue.Front()
		queue.Remove(node)

		qNode := node.Value.(QNode)

		left, right, sum := qNode.TreeNode.Left, qNode.TreeNode.Right, qNode.Sum
		if nil == left && nil == right {
			totalSum += sum // 此处已经没有子树了，加上 sum
		} else {
			if nil != left {
				queue.PushBack(QNode{left, sum*10 + left.Val})
			}

			if nil != right {
				queue.PushBack(QNode{right, sum*10 + right.Val})
			}
		}
	}
	return totalSum
}
