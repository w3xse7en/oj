package problem0108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  *TreeNode
}{
	{[]int{-10, -3, 0, 5, 9},
		&TreeNode{Val: 0,
			Left: &TreeNode{Val: -3,
				Left: &TreeNode{Val: -10},
			},
			Right: &TreeNode{Val: 5,
				Right: &TreeNode{Val: 9},
			},
		},
	},
	// {[]int{1, 2, 3, 4, 5, 6, 7}, nil},

	// 可以有多个 testcase
}

func Test_sortedArrayToBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, sortedArrayToBST(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_sortedArrayToBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortedArrayToBST(tc.nums)
		}
	}
}
