package problem0061

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	k    int
	ans  *ListNode
}{
	//{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{Val: 5}}}}}, 2, &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{Val: 3}}}}}},
	//{&ListNode{0, &ListNode{1, &ListNode{Val: 2}}}, 4, &ListNode{2, &ListNode{0, &ListNode{Val: 1}}}},
	{&ListNode{Val: 1}, 4, &ListNode{Val: 1}},

	// 可以有多个 testcase
}

func Test_rotateRight(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, rotateRight(tc.head, tc.k), "输入:%v", tc)
	}
}

func Benchmark_rotateRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotateRight(tc.head, tc.k)
		}
	}
}
