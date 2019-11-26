package main

import (
	"fmt"
)

// Given a sorted linked list, delete all duplicates such that each element appear only once.
//
// Example 1:
//
// Input: 1->1->2
// Output: 1->2
// Example 2:
//
// Input: 1->1->2->3->3
// Output: 1->2->3

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func PrintNodeList(node *ListNode1) {
	if node != nil {
		fmt.Println(node.Val)
		PrintNodeList(node.Next)
	}
}

func deleteDuplicates(head *ListNode1) *ListNode1 {
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

func main() {
	node1 := deleteDuplicates(&ListNode1{1, &ListNode1{1, &ListNode1{2, nil}}}) // 1 1 2 -> 1 2
	PrintNodeList(node1)
}
