package problem0024

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	return loop(head, head.Next)
}
func loop(l1, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return l1
	}
	t1 := l2.Next
	l2.Next = l1
	var t2 *ListNode
	if t1 != nil {
		t2 = t1.Next
	}
	l1.Next = loop(t1, t2)
	return l2
}
