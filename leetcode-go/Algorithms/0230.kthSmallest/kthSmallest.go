package _230_kthSmallest

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var val, cnt int

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	val = -1
	cnt = k
	dfs(root)
	return val
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	cnt--
	if cnt == 0 {
		val = root.Val
	}
	if cnt < 0 {
		return
	}
	dfs(root.Right)
	return
}
