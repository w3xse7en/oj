package problem0002

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tmp int
	h1, h2, l, h := l1, l2, &ListNode{}, &ListNode{}
	for l1 != nil || l2 != nil {
		a := 0
		if l1 != nil {
			a = l1.Val
		}
		b := 0
		if l2 != nil {
			b = l2.Val
		}
		sum := a + b + tmp
		if sum >= 10 {
			tmp = 1
			sum -= 10
		} else {
			tmp = 0
		}
		fmt.Println(sum, tmp)
		if l1 != nil {
			l = l1
			h = h1
			l1.Val = sum
			l1 = l1.Next
		}
		if l2 != nil {
			l = l2
			h = h2
			l2.Val = sum
			l2 = l2.Next
		}
	}
	if tmp != 0 {
		l.Next = &ListNode{Val: tmp}
	}
	return h
}
