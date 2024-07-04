package _095_deleteMiddle

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

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node, del, pre := head, head, head
	idx := 0
	for i := 0; node != nil; i++ {
		if idx < (i+1)/2 {
			pre = del
			del = del.Next
			idx++
		}
		node = node.Next
	}
	if idx == 0 {
		return nil
	}
	pre.Next = del.Next
	return head
}
