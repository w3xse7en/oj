package problem0160

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	for a, b := headA, headB; ; {
		if a == nil && b != nil {
			a = headB
			b = b.Next
		} else if a != nil && b == nil {
			b = headA
			a = a.Next
		} else if a != nil && b != nil && a == b {
			return a
		} else if a != nil && b != nil {
			a = a.Next
			b = b.Next
		} else if a == nil && b == nil {
			return nil
		}
	}
}
