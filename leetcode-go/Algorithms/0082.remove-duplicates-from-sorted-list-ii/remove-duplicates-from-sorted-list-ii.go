package problem0082

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	for head != nil {
		next, isRootDup := nextNoDupNode(head, false)
		if isRootDup {
			head = next
		} else {
			break
		}
	}
	node := head
	for node != nil {
		next, isRootDup := nextNoDupNode(node.Next, false)
		if isRootDup {
			node.Next = next
		} else {
			node = node.Next
		}
	}
	return head
}

func nextNoDupNode(node *ListNode, isRootDup bool) (*ListNode, bool) {
	if node == nil {
		return node, isRootDup
	}
	if node.Next != nil && node.Val == node.Next.Val {
		return nextNoDupNode(node.Next, true)
	}
	return node.Next, isRootDup
}
