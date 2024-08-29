package _023_mergeKLists

import (
	"math"
	"sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 2 {
		var l1, l2 *ListNode
		if len(lists) >= 1 && lists[0] != nil {
			l1 = lists[0]
		}
		if len(lists) >= 2 && lists[1] != nil {
			l2 = lists[1]
		}
		return mergeTwo(l1, l2)
	}
	l, r := 0, len(lists)
	mid := (l + r) / 2
	l1 := mergeKLists(lists[l:mid])
	l2 := mergeKLists(lists[mid:r])
	return mergeTwo(l1, l2)
}
func mergeTwo(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwo(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwo(list1, list2.Next)
	return list2
}

func linkNextNode(lists []*ListNode) *ListNode {
	sort.Slice(lists, func(i, j int) bool {
		a, b := math.MaxInt, math.MaxInt
		if lists[i] != nil {
			a = lists[i].Val
		}
		if lists[j] != nil {
			b = lists[j].Val
		}
		return a < b
	})
	if len(lists) == 0 {
		return nil
	}
	node := lists[0]
	if node == nil {
		return nil
	}
	lists[0] = lists[0].Next
	node.Next = linkNextNode(lists)
	return node
}
