package problem0112

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root != nil {
		if root.Left == nil && root.Right == nil && sum-root.Val == 0 {
			return true
		}
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
	return false
}
