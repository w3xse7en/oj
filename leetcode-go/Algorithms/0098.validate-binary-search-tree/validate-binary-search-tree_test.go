package problem0098

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  bool
}{
	//{root: nil, ans: true},
	{root: &TreeNode{Val: 1}, ans: false},
	//{root: &TreeNode{Val: 5,
	//	Left: &TreeNode{Val: 1},
	//	Right: &TreeNode{Val: 4,
	//		Left:  &TreeNode{Val: 3},
	//		Right: &TreeNode{Val: 6},
	//	},
	//}, ans: false},

	// 可以有多个 testcase
}

func Test_isValidBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isValidBST(tc.root), "输入:%v", tc)
	}
}

func Benchmark_isValidBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isValidBST(tc.root)
		}
	}
}
