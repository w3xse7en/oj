package problem0019

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	p1, p2, p3 := head, head, head
	for i, j := 0, 0; p1 != nil; {
		p1 = p1.Next
		if i == n {
			p2 = p2.Next
		} else {
			i++
		}
		if j == n+1 {
			p3 = p3.Next
		} else {
			j++
		}
	}
	if p2 == head {
		return p2.Next
	}
	p3.Next = p2.Next
	return head
}
