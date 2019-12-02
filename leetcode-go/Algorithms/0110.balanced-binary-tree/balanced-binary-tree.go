package problem0110

import (
	"github.com/w3xse7en/oj/leetcode-go/kit"
)

type TreeNode = kit.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ok, _ := problem0110LoopTree(root, 0)
	return ok
}

func problem0110LoopTree(root *TreeNode, dp int) (bool, int) {
	if root != nil {
		leftok, leftH := problem0110LoopTree(root.Left, dp+1)
		rightok, rightH := problem0110LoopTree(root.Right, dp+1)
		if !(leftok && rightok) {
			return false, 0
		}
		// fmt.Println(root.Val, dp, leftH, rightH)
		if leftH < rightH {
			leftH, rightH = rightH, leftH
		}
		if leftH-rightH > 1 {
			return false, 0
		}
		return true, leftH
	}
	return true, dp
}
