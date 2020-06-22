package problem0033

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    int
}{

	{[]int{}, 7, -1},
	{[]int{8, 7}, 7, 1},
	{[]int{8, 7}, 9, -1},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 9, -1},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 8, 7},
	{[]int{7, 8, 1, 2, 3, 4, 5, 6}, 8, 1},
	{[]int{8}, 8, 0},
	{[]int{8, 7}, 8, 0},
	{[]int{7, 8}, 8, 1},
	{[]int{7, 8, 1}, 8, 1},
	{[]int{1, 8, 7}, 8, 1},
	{[]int{8, 1, 2, 3, 4, 5, 6, 7}, 8, 0},
	{[]int{5, 6, 7, 8, 1, 2, 3, 4}, 8, 3},

	// 可以有多个 testcase
}

func Test_search(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, search(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_search(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			search(tc.nums, tc.target)
		}
	}
}
