package problem00222

import (
	. "github/veezhang/go-algorithm/leetcode/common/tree"
	"sort"
)

func countNodes1(root *TreeNode) int {
	var fn func(*TreeNode) int
	fn = func(root *TreeNode) int {
		if nil == root {
			return 0
		}

		return fn(root.Left) + fn(root.Right) + 1
	}
	return fn(root)
}

func countNodes2(root *TreeNode) int {
	getDepth := func(node *TreeNode) (depth int) {
		for ; nil != node; node = node.Left {
			depth++
		}
		return
	}
	var fn func(*TreeNode) int
	fn = func(root *TreeNode) int {
		if nil == root {
			return 0
		}
		lDepth, rDepth := getDepth(root.Left), getDepth(root.Right)
		if lDepth == rDepth {
			// 如果左边深度 == 右边深度，则左边一定是满二叉树，满二叉树个数为： 1<<lDepth - 1 ，再加上根节点
			return 1<<lDepth + fn(root.Right)
		} else {
			// 否则左边深度一定大于右边深度，右边一定为满二叉树，满二叉树个数为： 1<<lDepth - 1 ，再加上根节点
			return 1<<rDepth + fn(root.Left)
		}
	}

	return fn(root)
}

func countNodes3(root *TreeNode) int {
	if nil == root {
		return 0
	}
	// 树的深度
	depth := 0
	for node := root; nil != node; node = node.Left {
		depth++
	}

	// 把所有节点都从 1 开始标上下标
	// 则最后一排的下标为：[1 << (depth - 1), ... , 1 << depth - 1] ，也就是 [1 << (depth - 1), ... , 1 << depth)
	// 比如， depth = 3 ，最后一层可能的数据为 [4, 5, 6, 7]
	//     1
	//    / \
	//   2   3
	//  / \  /
	// 4  5 6
	// 把下标看成二进制，可以根据节点号计算其位置， 0 往左边移动， 1 往右边移动，
	// 比如 5(101)，除去最高位，得到 01 ，先左移动一下，再右移动一下就到 5 了
	// 最后再判断叶子节点是否为 nil

	// 最深那一排节点的第一个元素
	firstOfDeepest := 1 << (depth - 1)
	// 判断节点是否存在
	isExists := func(i int) bool {
		// 最后一次一定有一个，然后前面的都存在
		if i <= firstOfDeepest {
			return true // 其实这里也不会运行到，
		}
		// 从最高后的后一位( 1 << (depth - 2) )开始，依次遍历
		node := root
		for b := 1 << (depth - 2); b != 0; b >>= 1 {
			if 0 == b&i {
				node = node.Left
			} else {
				node = node.Right
			}
		}

		return nil != node
	}

	// func Search(n int, f func(int) bool) int {...}
	// Search 返回 [0,n) 中最小满足 f(i) = true 的 i
	// 查找最后一排中第一个节点为 nil 的
	// 在 [firstOfDeepest, 1<<depth) 中找
	firstNil := sort.Search(1<<depth-firstOfDeepest, func(i int) bool {
		return !isExists(i + firstOfDeepest)
	}) + firstOfDeepest

	// 返回，firstNil 为第一个为 nil 的，所以这里需要 - 1
	return firstNil - 1
}
