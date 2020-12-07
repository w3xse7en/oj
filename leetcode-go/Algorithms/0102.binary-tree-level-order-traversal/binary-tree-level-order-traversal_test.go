package problem0102

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  [][]int
}{

	{root: &TreeNode{Val: 5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}, ans: [][]int{}},

	// 可以有多个 testcase
}

func Test_list(t *testing.T) {
	a := [][]int{{1}}
	add(a)
	fmt.Println(a)
}
func add(a [][]int) {
	a[0] = append(a[0], 2)
	a[1] = []int{1}
}
func Test_levelOrder(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, levelOrder(tc.root), "输入:%v", tc)
	}
}

func Benchmark_levelOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			levelOrder(tc.root)
		}
	}
}
