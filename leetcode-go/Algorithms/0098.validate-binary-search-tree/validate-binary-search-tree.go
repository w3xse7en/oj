package problem0098

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

var preValue *int
var isValid bool

func isValidBST(root *TreeNode) bool {
	isValid = true
	preValue = nil
	loop(root)
	return isValid
}

func loop(root *TreeNode) {
	if isValid && root != nil {
		loop(root.Left)
		if preValue != nil && *preValue >= root.Val {
			isValid = false
		}
		preValue = &root.Val
		loop(root.Right)
	}
}
