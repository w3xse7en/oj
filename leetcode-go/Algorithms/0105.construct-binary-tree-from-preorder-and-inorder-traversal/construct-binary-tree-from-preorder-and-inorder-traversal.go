package problem0105

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	f(root, preorder[0], preorder, inorder)
	return root
}

func f(root *TreeNode, val int, preorder, inorder []int) {
	idx := in(val, inorder)
	if idx >= 0 {
		addNewNode(root, preorder[1:idx+1], preorder[idx+1:], inorder[:idx], inorder[idx+1:])
	}
}

func addNewNode(root *TreeNode, leftPre, rightPre, leftIn, rightIn []int) {
	if len(leftPre) == 0 && len(rightPre) == 0 {
		return
	}

	if len(leftPre) != 0 {
		pre := leftPre
		val := pre[0]
		newNode := &TreeNode{Val: val}
		root.Left = newNode
		f(newNode, val, pre, leftIn)
		f(newNode, val, pre, rightIn)
	}
	if len(rightPre) != 0 {
		pre := rightPre
		val := pre[0]
		newNode := &TreeNode{Val: val}
		root.Right = newNode
		f(newNode, val, pre, leftIn)
		f(newNode, val, pre, rightIn)
	}
}

func in(target int, list []int) int {
	for k, v := range list {
		if target == v {
			return k
		}
	}
	return -1
}
