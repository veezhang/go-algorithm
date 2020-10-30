package problem00124

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64

	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if nil == root {
			return 0
		}
		// 计算左右边得收益，负收益被舍弃，所以 max(0, ...)
		leftGain := max(0, dfs(root.Left))
		rightGain := max(0, dfs(root.Right))

		// 收益为：当前节点 + 左节点 + 右节点
		// 更新下最大收益值，也就是最大路径
		maxSum = max(maxSum, root.Val+leftGain+rightGain)

		// 返回为左节点， 右节点中收益较大的
		return root.Val + max(leftGain, rightGain)
	}

	dfs(root)

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
