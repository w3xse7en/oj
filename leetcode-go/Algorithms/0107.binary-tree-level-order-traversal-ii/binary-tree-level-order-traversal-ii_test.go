package problem0107

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  [][]int
}{
	{
		&TreeNode{Val: 3,
			Left: &TreeNode{Val: 9},

			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		[][]int{{15, 7}, {9, 20}, {3}},
	},
	{
		nil,
		[][]int{},
	},

	// 可以有多个 testcase
}

func Test_levelOrderBottom(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, levelOrderBottom(tc.root), "输入:%v", tc)
	}
}

func Benchmark_levelOrderBottom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			levelOrderBottom(tc.root)
		}
	}
}
