package problem0109

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode
type TreeNode = kit.TreeNode

func sortedListToBST(head *ListNode) *TreeNode {
	var list []int
	for ; head != nil; head = head.Next {
		list = append(list, head.Val)
	}
	var root *TreeNode

	return loop(list, root)
}
func addNode(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if v > root.Val {
		if root.Right == nil {
			root.Right = &TreeNode{Val: v}
		} else {
			addNode(root.Right, v)
		}
	} else {
		if root.Left == nil {
			root.Left = &TreeNode{Val: v}
		} else {
			addNode(root.Left, v)
		}
	}
	return root
}

func loop(list []int, root *TreeNode) *TreeNode {
	l := len(list)
	if l == 0 {
		return root
	}
	if l == 1 {
		return addNode(root, list[0])
	}
	mid := l / 2
	root = addNode(root, list[mid])
	loop(list[:mid], root)
	loop(list[mid+1:], root)

	return root
}
