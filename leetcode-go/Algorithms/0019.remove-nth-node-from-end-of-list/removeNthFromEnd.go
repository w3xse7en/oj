package problem0019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := loop(head, []*ListNode{})
	if len(list) == 0 {
		return nil
	}
	step := len(list) - n - 1
	if step < 0 {
		return head.Next
	}
	if list[step].Next != nil {
		list[step].Next = list[step].Next.Next
	}
	return head
}
func loop(node *ListNode, list []*ListNode) []*ListNode {
	if node == nil {
		return list
	}
	list = append(list, node)
	node = node.Next
	return loop(node, list)
}
