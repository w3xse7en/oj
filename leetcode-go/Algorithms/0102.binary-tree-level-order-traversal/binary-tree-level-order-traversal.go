package problem0102

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	loop(root, &result, 0)
	return result
}
func loop(root *TreeNode, result *[][]int, high int) {
	if root != nil {
		if len(*result) == high {
			*result = append(*result, []int{root.Val})
		} else {
			(*result)[high] = append((*result)[high], root.Val)
		}
		loop(root.Left, result, high+1)
		loop(root.Right, result, high+1)
	}
}
