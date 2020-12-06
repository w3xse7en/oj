package problem0094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []int
}{

	// 可以有多个 testcase
}

func Test_inorderTraversal(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, inorderTraversal(tc.root), "输入:%v", tc)
	}
}

func Benchmark_inorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			inorderTraversal(tc.root)
		}
	}
}
