package problem0094

import (
	"fmt"
	"github.com/w3xse7en/oj/leetcode-go/kit"
)

type TreeNode = kit.TreeNode

var vals []int

func inorderTraversal(root *TreeNode) []int {
	vals = make([]int, 0)
	loop(root)
	return vals
}

func loop(root *TreeNode) {
	if root != nil {
		fmt.Println(1, root.Val)
		loop(root.Left)
		fmt.Println(2, root.Val)
		vals = append(vals, root.Val)
		loop(root.Right)
		fmt.Println(3, root.Val)
	}
}
