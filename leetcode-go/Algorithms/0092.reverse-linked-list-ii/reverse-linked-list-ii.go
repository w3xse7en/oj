package problem0092

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	var prem, mnode, nnode, aftern, tail *ListNode
	for i, node := 1, head; node != nil && i <= n+1; i++ {
		next := node.Next
		if i == m-1 {
			prem = node
		}
		if i == m {
			mnode = node
			mnode.Next = nil
			tail = node
		}
		if i == n {
			nnode = node
			nnode.Next = nil
		}
		if i == n+1 {
			aftern = node
		}
		if i >= m+1 && i <= n {
			node.Next = tail
			tail = node
		}
		node = next
	}
	if mnode != nil {
		mnode.Next = aftern
	}
	if prem != nil {
		prem.Next = nnode
	}
	if m == 1 {
		return nnode
	}
	return head
}
