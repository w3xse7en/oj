package _669_mergeInBetween

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var l1a, l1b *ListNode
	l1 := list1
	for i := 0; l1 != nil; i++ {
		if i == a-1 {
			l1a = l1
		}
		if i == b {
			l1b = l1.Next
			l1.Next = nil
			break
		}
		l1 = l1.Next
	}
	l1a.Next = list2
	if l1b != nil {
		for list2.Next != nil {
			list2 = list2.Next
		}
		list2.Next = l1b
	}
	return list1
}
