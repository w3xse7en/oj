package problem0082

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	head *ListNode
	ans  *ListNode
}{
	{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}, &ListNode{Val: 3}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}, nil},
	{&ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}, &ListNode{Val: 1}},
	{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}},
	{nil, nil},
	// 可以有多个 testcase
}

func Test_nextNoDupNode(t *testing.T) {
	fmt.Println(nextNoDupNode(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}, false))
	fmt.Println(nextNoDupNode(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}, false))
}

func Test_deleteDuplicates(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, deleteDuplicates(tc.head), "输入:%v", tc)
	}
}

func Benchmark_deleteDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deleteDuplicates(tc.head)
		}
	}
}
