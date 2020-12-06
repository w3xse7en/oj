package problem0092

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	m    int
	n    int
	ans  *ListNode
}{
	{nil, 1, 1, nil},
	{&ListNode{Val: 1}, 1, 1, &ListNode{Val: 1}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, 1, 2, &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, 1, 3, &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, 2, 3, &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, 1, 2, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, 2, 2, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},

	// 可以有多个 testcase
}

func Test_reverseBetween(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, reverseBetween(tc.head, tc.m, tc.n), "输入:%v", tc)
	}
}

func Benchmark_reverseBetween(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseBetween(tc.head, tc.m, tc.n)
		}
	}
}
