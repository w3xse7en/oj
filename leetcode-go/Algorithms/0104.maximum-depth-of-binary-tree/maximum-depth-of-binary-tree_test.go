package problem0104

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  int
}{
	{
		&TreeNode{Val: 1,
			Left: &TreeNode{Val: 2,
				Left: &TreeNode{Val: 3}},

			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 3},
			},
		},
		3,
	},
	{
		&TreeNode{Val: 1},
		1,
	},
	{
		nil,
		0,
	},
	// {
	// 	&TreeNode{Val: 1,
	// 		Left: &TreeNode{Val: 1,
	// 			Left: &TreeNode{Val: 3}},
	//
	// 		Right: &TreeNode{Val: 2,
	// 			Right: &TreeNode{Val: 3},
	// 		},
	// 	},
	// 	3,
	// },
	// {
	// 	&TreeNode{Val: 1,
	// 		Left: &TreeNode{Val: 1,
	// 			Left:  &TreeNode{Val: 3},
	// 			Right: &TreeNode{Val: 4},
	// 		},
	//
	// 		Right: &TreeNode{Val: 1,
	// 			Left:  &TreeNode{Val: 4},
	// 			Right: &TreeNode{Val: 3},
	// 		},
	// 	},
	// 	3,
	// },
	// {
	// 	&TreeNode{Val: 1,
	// 		Left: &TreeNode{Val: 2,
	// 			Left:  nil,
	// 			Right: &TreeNode{Val: 3},
	// 		},
	//
	// 		Right: &TreeNode{Val: 2,
	// 			Left:  nil,
	// 			Right: &TreeNode{Val: 3},
	// 		},
	// 	},
	// 	3,
	// },

	// 可以有多个 testcase
}

func Test_maxDepth(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxDepth(tc.root), "输入:%v", tc)
	}
}

func Benchmark_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxDepth(tc.root)
		}
	}
}
