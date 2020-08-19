package problem0061

import (
	"github.com/w3xse7en/oj/leetcode-go/kit"
)

type ListNode = kit.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	node := head
	length := 1
	for ; node.Next != nil; node = node.Next {
		length++
	}
	node.Next = head
	k %= length
	for i := 1; i < length-k; i++ {
		head = head.Next
	}
	newHead := head.Next
	head.Next = nil
	return newHead
}
