package problem0141

import "github.com/w3xse7en/oj/leetcode-go/kit"

type ListNode = kit.ListNode

func hasCycle(head *ListNode) bool {
	for h1, h2 := head, head; h1 != nil && h2 != nil; {
		h2 = h2.Next
		if h2 == h1 {
			return true
		} else if h2 != nil {
			h2 = h2.Next
		}
		h1 = h1.Next
	}
	return false
}

func hasCycleOrigin(head *ListNode) bool {
	exist := make(map[*ListNode]bool)
	for head != nil {
		if exist[head] {
			return true
		} else {
			exist[head] = true
		}
		head = head.Next
	}
	return false
}
