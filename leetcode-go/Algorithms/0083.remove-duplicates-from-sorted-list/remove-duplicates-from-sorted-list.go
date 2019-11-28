package problem0083

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	for node := head; node != nil; {
		if node.Next != nil {
			if node.Val == node.Next.Val {
				node.Next = node.Next.Next
				continue
			}
		}
		node = node.Next
	}
	return head
}
