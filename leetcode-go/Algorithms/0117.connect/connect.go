package _117_connect

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

var nodeList []*Node

func connect(root *Node) *Node {
	nodeList = make([]*Node, 6000)
	dfs(root, 0)
	return root
}

func dfs(node *Node, deep int) {
	if node == nil {
		return
	}
	if nodeList[deep] != nil {
		nodeList[deep].Next = node
	}
	nodeList[deep] = node
	dfs(node.Left, deep+1)
	dfs(node.Right, deep+1)
}
