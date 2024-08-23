package _138_copyRandomList

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	existOldNode := map[*Node]*Node{}
	getNewNode := func(old *Node) *Node {
		if old == nil {
			return nil
		}
		if n, ok := existOldNode[old]; ok {
			return n
		}
		n := &Node{Val: old.Val}
		existOldNode[old] = n
		return n
	}
	old := head
	newHead := getNewNode(old)
	for old != nil {
		newNode := getNewNode(old)
		nextNewNode := getNewNode(old.Next)
		randomNewNode := getNewNode(old.Random)
		newNode.Next = nextNewNode
		newNode.Random = randomNewNode
		old = old.Next
	}
	return newHead
}
