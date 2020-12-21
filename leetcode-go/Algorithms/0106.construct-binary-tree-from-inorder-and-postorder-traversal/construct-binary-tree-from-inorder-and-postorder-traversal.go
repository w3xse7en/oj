package problem0106

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	f(root, postorder[len(postorder)-1], postorder, inorder)
	return root
}

// in 		9,3,15,20,7
// post 	9,15,7,20,3
// reverse 	3,20,7,15,9

func f(root *TreeNode, val int, postorder, inorder []int) {
	idx := in(val, inorder)
	if idx >= 0 {
		addNewNode(root, postorder[:idx], postorder[idx:len(postorder)-1], inorder[:idx], inorder[idx+1:])
	}
}

func addNewNode(root *TreeNode, leftPost, rightPost, leftIn, rightIn []int) {
	if len(leftPost) == 0 && len(rightPost) == 0 {
		return
	}

	if len(leftPost) != 0 {
		post := leftPost
		val := post[len(post)-1]
		newNode := &TreeNode{Val: val}
		root.Left = newNode
		f(newNode, val, post, leftIn)
		f(newNode, val, post, rightIn)
	}
	if len(rightPost) != 0 {
		post := rightPost
		val := post[len(post)-1]
		newNode := &TreeNode{Val: val}
		root.Right = newNode
		f(newNode, val, post, leftIn)
		f(newNode, val, post, rightIn)
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
