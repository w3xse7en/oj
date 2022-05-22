package problem6057

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	_, _, match := dfs(root)
	return match
}

func dfs(node *TreeNode) (total int, cnt int, match int) {
	if node == nil {
		return 0, 0, 0
	}
	lt, lc, lm := dfs(node.Left)
	rt, rc, rm := dfs(node.Right)
	m := lm + rm
	if (node.Val+lt+rt)/(lc+rc+1) == node.Val {
		m++
	}
	fmt.Println(lt, lc, rt, rc, m, node.Val)
	return lt + rt + node.Val, lc + rc + 1, m
}
