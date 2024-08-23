package _543_diameterOfBinaryTree

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

var mx int

func diameterOfBinaryTree(root *TreeNode) int {
	mx = 0
	dfs(root)
	return mx
}

func dfs(node *TreeNode) int {
	if node == nil {
		return -1
	}
	l := dfs(node.Left) + 1
	r := dfs(node.Right) + 1
	mx = max(mx, l+r)
	return max(l, r)
}
