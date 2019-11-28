package problem0101

import "github.com/w3xse7en/oj/leetcode-go/kit"

type TreeNode = kit.TreeNode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return problem0101LoopTree(root.Left, root.Right)
}
func problem0101LoopTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return problem0101LoopTree(left.Left, right.Right) && problem0101LoopTree(left.Right, right.Left)
}

func isSymmetricGoroutine(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nodeLeft := make(chan *TreeNode)
	nodeRight := make(chan *TreeNode)
	endLeft := make(chan bool)
	endRight := make(chan bool)
	go func() {
		problem0101LoopTreeLeftF(root.Left, nodeLeft)
		endLeft <- true
	}()
	go func() {
		problem0101LoopTreeRightF(root.Right, nodeRight)
		endRight <- true
	}()
	var el, er bool
	for {
		var l, r *TreeNode
		select {
		case l = <-nodeLeft:
		case el = <-endLeft:
		}
		select {
		case r = <-nodeRight:
		case er = <-endRight:
		}
		if (l == nil && r != nil) || (l != nil && r == nil) || (l != nil && r != nil && l.Val != r.Val) {
			return false
		}
		if el && er {
			return true
		}
	}
}

func problem0101LoopTreeLeftF(node *TreeNode, c chan *TreeNode) {
	c <- node
	if node != nil {
		problem0101LoopTreeLeftF(node.Left, c)
		problem0101LoopTreeLeftF(node.Right, c)
	}
}

func problem0101LoopTreeRightF(node *TreeNode, c chan *TreeNode) {
	c <- node
	if node != nil {
		problem0101LoopTreeRightF(node.Right, c)
		problem0101LoopTreeRightF(node.Left, c)
	}
}
