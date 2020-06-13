package problem0024

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  *ListNode
}{

	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2}}, &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
	{&ListNode{Val: 1}, &ListNode{Val: 1}},
	{&ListNode{}, &ListNode{}},

	// 可以有多个 testcase
}

func Test_swapPairs(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, swapPairs(tc.head), "输入:%v", tc)
	}
}

func Benchmark_swapPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			swapPairs(tc.head)
		}
	}
}
