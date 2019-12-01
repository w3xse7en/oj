package problem0108

import (
	"github.com/w3xse7en/oj/leetcode-go/kit"
)

type TreeNode = kit.TreeNode

func sortedArrayToBSTBetter(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	half := len(nums) / 2
	return &TreeNode{
		Val:   nums[half],
		Left:  sortedArrayToBSTBetter(nums[:half]),
		Right: sortedArrayToBSTBetter(nums[half+1:]),
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := new(TreeNode)
	problem0108LoopInt(nums, root)
	return root
}

func problem0108LoopInt(nums []int, root *TreeNode) {
	if len(nums) != 0 {
		half := len(nums) / 2
		root.Val = nums[half]

		nLeft := nums[:half]
		if len(nLeft) != 0 {
			root.Left = new(TreeNode)
		}
		problem0108LoopInt(nLeft, root.Left)
		nRight := nums[half+1:]
		if len(nRight) != 0 {
			root.Right = new(TreeNode)
		}
		problem0108LoopInt(nRight, root.Right)
	}
}
