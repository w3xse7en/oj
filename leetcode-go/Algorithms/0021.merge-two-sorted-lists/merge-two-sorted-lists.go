package problem0021

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	apd(l1, l2)

	return l1
}

func apd(root *ListNode, insert *ListNode) {
	if insert == nil {
		return
	}
	if root.Next == nil {
		root.Next = insert
		return
	}
	if root.Next.Val < insert.Val {
		apd(root.Next, insert)
		return
	}
	nextInsert := root.Next
	root.Next = insert
	apd(insert, nextInsert)
}
