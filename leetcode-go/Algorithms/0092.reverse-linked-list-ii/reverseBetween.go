package problem0092

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	node, l, r := head, head, head
	ll, rr := head, head
	for i := 0; node != nil; i++ {
		r = r.Next
		ll = l
		l = l.Next
		if i == left-1 {

		}
		node = node.Next
	}

	return nil
}
