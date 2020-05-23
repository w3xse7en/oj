package problem0203

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	for node := head; node != nil; {
		if node.Next != nil && node.Next.Val == val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	if head != nil && head.Val == val {
		return head.Next
	}
	return head
}
