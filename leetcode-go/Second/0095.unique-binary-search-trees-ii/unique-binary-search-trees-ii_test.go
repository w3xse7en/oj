package problem0095

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans []*TreeNode
}{
	{7, nil},
	// 可以有多个 testcase
}

func Test_addnode(t *testing.T) {
	root := &TreeNode{Val: 1}
	addNode(2, root)
	fmt.Printf("%+v", root)
}
func Test_ref(t *testing.T) {
	root := TreeNode{Val: 1}
	buildleft(&root)
	buildright(&root)
	fmt.Printf("%+v", root)
}
func buildleft(root *TreeNode) {
	root.Left = &TreeNode{Val: 2}
}
func buildright(root *TreeNode) {
	root.Right = &TreeNode{Val: 3}
}
func Test_str(t *testing.T) {
	s := "1234"
	s = s[:3] + "0" + s[4:]
	add(s)
	fmt.Println(s)
}
func add(s string) {
	s += "12344444"
}

func Test_list(t *testing.T) {
	list := []int{1}
	list = append(list[:0], list[1:]...)
	fmt.Println(list)
}
func Test_generateTrees(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, generateTrees(tc.n), "输入:%v", tc)
	}
}

func Benchmark_generateTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			generateTrees(tc.n)
		}
	}
}
