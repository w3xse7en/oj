package problem0111

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDp := new(int)
	*minDp = 1000000
	problem0111LoopTree(root, minDp, 1)
	return *minDp
}

func problem0111LoopTree(root *TreeNode, minDp *int, dp int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			if *minDp > dp {
				*minDp = dp
			}
		}
		problem0111LoopTree(root.Left, minDp, dp+1)
		problem0111LoopTree(root.Right, minDp, dp+1)
	}
}

// func minDepth(root *TreeNode) int {
// 	dp := 0
// 	que, nextQue := make([]*TreeNode, 0), make([]*TreeNode, 0)
// 	if root != nil {
// 		que = append(que, root)
// 	}
// 	for ; len(que) != 0; {
// 		dp++
// 		for _, v := range que {
// 			if v.Left == nil && v.Right == nil {
// 				return dp
// 			}
// 			if v.Left != nil {
// 				nextQue = append(nextQue, v.Left)
// 			}
// 			if v.Right != nil {
// 				nextQue = append(nextQue, v.Right)
// 			}
// 		}
// 		que = nextQue
// 	}
// 	return dp
// }
