package problem0206

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func reverseList(head *ListNode) *ListNode {
	var rHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = rHead
		rHead = head
		head = tmp
	}
	return rHead
}
