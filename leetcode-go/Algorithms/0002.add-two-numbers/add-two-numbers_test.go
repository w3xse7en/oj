package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	l1  *ListNode
	l2  *ListNode
	ans *ListNode
}{
	{
		l1: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		l2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		ans: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  8,
					Next: nil,
				},
			},
		},
	},

	{&ListNode{Val: 5}, &ListNode{Val: 5}, &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
	{nil, nil, nil},
	{&ListNode{Val: 0}, nil, &ListNode{Val: 0}},
	{nil, &ListNode{Val: 1}, &ListNode{Val: 1}},
	{nil, &ListNode{Val: 0}, &ListNode{Val: 0}},
	{
		l1: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
		l2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		ans: &ListNode{
			Val: 7,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	},
	// 可以有多个 testcase
}

func Test_addTwoNumbers(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, addTwoNumbers(tc.l1, tc.l2), "输入:%v", tc)
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			addTwoNumbers(tc.l1, tc.l2)
		}
	}
}
