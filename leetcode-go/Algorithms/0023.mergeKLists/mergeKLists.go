package _023_mergeKLists

import (
	"math"
	"sort"
)

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

func mergeKLists1(lists []*ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for len(lists) != 0 {
		sort.Slice(lists, func(i, j int) bool {
			a, b := math.MinInt, math.MinInt
			if lists[i] != nil {
				a = lists[i].Val
			}
			if lists[j] != nil {
				b = lists[j].Val
			}
			return a < b
		})
		if lists[0] == nil {
			lists = lists[1:]
			continue
		}
		node.Next = lists[0]
		node = node.Next
		lists[0] = lists[0].Next
	}
	return head.Next
}
