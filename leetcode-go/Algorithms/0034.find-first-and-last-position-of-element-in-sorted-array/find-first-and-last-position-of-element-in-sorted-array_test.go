package problem0034

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    []int
}{
	{[]int{1}, 1, []int{0, 0}},
	{[]int{1, 2}, 2, []int{1, 1}},
	{[]int{1, 2}, 3, []int{-1, -1}},
	{[]int{1}, 3, []int{-1, -1}},
	{[]int{}, 3, []int{-1, -1}},
	{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	{[]int{2, 2, 2, 2, 2, 2, 2}, 2, []int{0, 6}},
	{[]int{1, 2, 2, 2, 2, 2, 3}, 2, []int{1, 5}},

	// 可以有多个 testcase
}

func Test_searchRange(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, searchRange(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_searchRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			searchRange(tc.nums, tc.target)
		}
	}
}
