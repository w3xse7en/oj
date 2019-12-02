package problem0110

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  bool
}{
	{
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		},
		false,
	},
	// {
	// 	&TreeNode{Val: 1,
	// 		Right: &TreeNode{Val: 2,
	// 			Right: &TreeNode{Val: 3},
	// 		},
	// 	},
	// 	false,
	// },

	// 可以有多个 testcase
}

func Test_isBalanced(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isBalanced(tc.root), "输入:%v", tc)
	}
}

func Benchmark_isBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isBalanced(tc.root)
		}
	}
}
