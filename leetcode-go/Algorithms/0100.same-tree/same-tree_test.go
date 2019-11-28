package problem0100

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	p   *TreeNode
	q   *TreeNode
	ans bool
}{
	{
		&TreeNode{Left: &TreeNode{Val: 2}, Val: 1, Right: &TreeNode{Val: 3}},
		&TreeNode{Left: &TreeNode{Val: 2}, Val: 1, Right: &TreeNode{Val: 3}},
		true,
	},
	{
		&TreeNode{Left: &TreeNode{Val: 2}, Val: 1, Right: &TreeNode{Val: 3}},
		&TreeNode{Left: &TreeNode{Val: 3}, Val: 1, Right: &TreeNode{Val: 2}},
		false,
	},
	{
		nil,
		&TreeNode{Left: &TreeNode{Val: 3}, Val: 1, Right: &TreeNode{Val: 2}},
		false,
	},
	{
		nil,
		nil,
		true,
	},
	{
		&TreeNode{Left: &TreeNode{Val: 0}, Val: 0, Right: &TreeNode{Val: 0}},
		&TreeNode{Left: &TreeNode{Val: 0}, Val: 0, Right: &TreeNode{Val: 1}},
		false,
	},

	// 可以有多个 testcase
}

func Test_isSameTree(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isSameTree(tc.p, tc.q), "输入:%v", tc)
	}
}

func Benchmark_isSameTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSameTree(tc.p, tc.q)
		}
	}
}
