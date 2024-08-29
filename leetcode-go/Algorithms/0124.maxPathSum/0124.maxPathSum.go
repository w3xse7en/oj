package _124_maxPathSum

import "math"

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

func maxPathSum(root *TreeNode) int {
	mx = math.MinInt
	dfs(root)
	return mx
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lv := dfs(node.Left)
	rv := dfs(node.Right)
	nodeMx := max(node.Val, node.Val+lv, node.Val+rv)
	//fmt.Println(node.Val, lv, rv, nodeMx)
	mx = max(mx, node.Val+lv+rv, node.Val, node.Val+lv, node.Val+rv)
	return nodeMx
}
