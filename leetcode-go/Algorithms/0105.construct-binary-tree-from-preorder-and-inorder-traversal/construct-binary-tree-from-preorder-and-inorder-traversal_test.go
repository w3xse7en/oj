package problem0105

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	preorder []int
	inorder  []int
	ans      *TreeNode
}{
	{[]int{1, 2, 4, 8, 5, 9, 3, 6, 7, 10}, []int{4, 8, 2, 5, 9, 1, 6, 3, 7, 10}, &TreeNode{}},
	//{[]int{3,9,20,15,7},[]int{9,3,15,20,7},&TreeNode{}},

	// 可以有多个 testcase
}

func Test_buildTree(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, buildTree(tc.preorder, tc.inorder), "输入:%v", tc)
	}
}

func Benchmark_buildTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			buildTree(tc.preorder, tc.inorder)
		}
	}
}
