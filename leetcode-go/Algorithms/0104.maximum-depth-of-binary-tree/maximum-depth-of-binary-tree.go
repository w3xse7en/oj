package problem0104

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func problem0104LoopTree(root *TreeNode, depth int, maxDp *int) {
	if root != nil {
		if depth > *maxDp {
			*maxDp = depth
		}
		problem0104LoopTree(root.Left, depth+1, maxDp)
		problem0104LoopTree(root.Right, depth+1, maxDp)
	}
}

func maxDepth(root *TreeNode) int {
	var maxDp int
	problem0104LoopTree(root, 1, &maxDp)
	return maxDp
}
