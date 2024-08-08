package problem0080

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{
	//{[]int{1,2},2},
	//{[]int{1,2,3},3},
	//{[]int{0,0,0,0,0},2},
	//{[]int{1,1,1,2,2,3},5},
	//{[]int{0,0,1,1,1,1,2,3,3},7},
	//{[]int{0,0,1,1,1,1,1,2,3,3},7},
	{[]int{0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3}, 8},
	//{[]int{0,0},2},
	//{[]int{0},1},
	//{[]int{},0},

	// 可以有多个 testcase
}

func Test_removeDuplicates1(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, removeDuplicates(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			removeDuplicates(tc.nums)
		}
	}
}
