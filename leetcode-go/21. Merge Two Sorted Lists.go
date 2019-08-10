package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	AHead := &ListNode{Val: 2}
	ANext1 := &ListNode{Val: 20}
	ANext2 := &ListNode{Val: 30}
	AHead.Next = ANext1
	ANext1.Next = ANext2

	BHead := &ListNode{Val: 10}
	BNext1 := &ListNode{Val: 13}
	BNext2 := &ListNode{Val: 33}
	BHead.Next = BNext1
	BNext1.Next = BNext2
	node := mergeTwoLists(AHead, BHead)
	pr(node)
}

func pr(node *ListNode) {
	if node != nil {
		fmt.Println(node.Val)
		pr(node.Next)
	}
}
