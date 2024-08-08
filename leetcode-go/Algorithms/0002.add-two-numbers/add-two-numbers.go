package problem0002

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	if l1 == nil && l2 == nil {
		return l1
	}

	var (
		up    int
		last  *ListNode
		merge bool
	)
	for n1, n2 := l1, l2; n1 != nil && n2 != nil; {
		sum := n1.Val + n2.Val + up
		if merge {
			sum = n1.Val + up
		}
		if sum >= 10 {
			sum -= 10
			up = 1
		} else {
			up = 0
		}
		n1.Val = sum
		n2.Val = sum
		if n1.Next == nil && n2.Next != nil {
			n1.Next = n2.Next
			merge = true
		}
		if n1.Next != nil && n2.Next == nil {
			n2.Next = n1.Next
			merge = true
		}
		last = n1
		n1 = n1.Next
		n2 = n2.Next
	}
	if last != nil && up > 0 {
		last.Next = &ListNode{Val: up}
	}
	return l1
}
