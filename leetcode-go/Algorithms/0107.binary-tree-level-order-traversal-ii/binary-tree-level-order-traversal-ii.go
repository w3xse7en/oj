package problem0107

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	problem0107TreeLoop(root, &result, 0)
	length := len(result)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func problem0107TreeLoop(root *TreeNode, result *[][]int, dep int) {
	if root != nil {
		if len(*result) <= dep {
			*result = append(*result, []int{root.Val})
		} else {
			(*result)[dep] = append((*result)[dep], root.Val)
		}
		problem0107TreeLoop(root.Left, result, dep+1)
		problem0107TreeLoop(root.Right, result, dep+1)
	}
}

func levelOrderBottomOrigin(root *TreeNode) [][]int {
	result := make([][]int, 0)
	que := make([]*TreeNode, 0)
	if root != nil {
		que = append(que, root)
	}
	for len(que) != 0 {
		tmpResult := make([]int, 0)
		tmpQue := make([]*TreeNode, 0)
		for _, v := range que {
			if v.Left != nil {
				tmpQue = append(tmpQue, v.Left)
			}
			if v.Right != nil {
				tmpQue = append(tmpQue, v.Right)
			}
			tmpResult = append(tmpResult, v.Val)
		}
		que = tmpQue
		result = append([][]int{tmpResult}, result...)
	}
	return result
}
