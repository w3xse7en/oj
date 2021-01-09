package problem0109

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  *TreeNode
}{
	{&ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}, nil},

	// 可以有多个 testcase
}

func Test_sortedListToBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, sortedListToBST(tc.head), "输入:%v", tc)
	}
}

func Benchmark_sortedListToBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortedListToBST(tc.head)
		}
	}
}
