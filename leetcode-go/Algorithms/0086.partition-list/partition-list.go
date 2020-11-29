package problem0086

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func partition(head *ListNode, x int) *ListNode {
	var lessXHead, lessXNode, grateXHead, grateXNode *ListNode
	for head != nil {
		node := head
		head = head.Next
		node.Next = nil
		if node.Val < x {
			if lessXNode == nil {
				lessXHead = node
				lessXNode = node
			} else {
				lessXNode.Next = node
				lessXNode = node
			}
		} else {
			if grateXNode == nil {
				grateXHead = node
				grateXNode = node
			} else {
				grateXNode.Next = node
				grateXNode = node
			}
		}
	}
	if lessXNode != nil {
		lessXNode.Next = grateXHead
		return lessXHead
	}
	if grateXNode != nil {
		return grateXHead
	}
	return head
}
