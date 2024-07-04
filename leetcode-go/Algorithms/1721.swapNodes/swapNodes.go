package _721_swapNodes

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	a, b, node := head, head, head.Next
	for i := 1; node != nil; i++ {
		if i < k {
			a = a.Next
		} else {
			b = b.Next
		}
		node = node.Next
	}
	a.Val, b.Val = b.Val, a.Val
	return head
}

func loop(node *ListNode, list []*ListNode) []*ListNode {
	if node == nil {
		return list
	}
	return loop(node.Next, append(list, node))
}
