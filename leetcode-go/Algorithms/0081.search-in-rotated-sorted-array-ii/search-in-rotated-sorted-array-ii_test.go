package problem0081

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    bool
}{
	{[]int{1, 3, 1, 1, 1}, 3, true},
	{[]int{1, 1, 1, 3, 1}, 3, true},
	{[]int{4, 5, 6, 7, 8, 9, 9, 0, 1, 2, 3}, 9, true},

	{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 2, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 5, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 6, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 1, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 2, true},
	{[]int{2, 5, 6, 0, 0, 1, 2}, -1, false},
	{[]int{2, 5, 6, 0, 0, 1, 2}, 9, false},
	//
	{[]int{1, 2, 3}, 9, false},
	{[]int{1, 2, 3}, -1, false},
	{[]int{1, 2, 3}, 1, true},
	{[]int{1, 2, 3}, 2, true},
	{[]int{1, 2, 3}, 3, true},
	{[]int{1}, 1, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 10, false},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, -1, false},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 9, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 8, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 7, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 6, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 6, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 6, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 6, true},
	{[]int{4, 5, 5, 6, 7, 7, 7, 7, 7, 7, 8, 9, 0, 0, 0, 0, 0, 1, 2, 3}, 6, true},
	{[]int{4, 5, 5, 6, 7, 7, 7, 7, 7, 7, 8, 9, 0, 0, 0, 0, 0, 1, 2, 3}, -1, false},
	{[]int{4, 5, 5, 6, 7, 7, 7, 7, 7, 7, 8, 9, 0, 0, 0, 0, 0, 1, 2, 3}, 10, false},
	{[]int{4, 5, 5, 6, 7, 7, 7, 7, 7, 7, 8, 9, 0, 0, 0, 0, 0, 1, 2, 3}, 7, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 5, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 4, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 3, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 2, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 1, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 9, true},
	{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, 0, true},
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

func Test_binarySearch(t *testing.T) {
	//fmt.Println(bs([]int{1, 2, 3, 4}, 5))
	//fmt.Println(bs([]int{1, 2, 3, 4}, 4))
	//fmt.Println(bs([]int{1, 2, 3, 4}, 3))
	//fmt.Println(bs([]int{1, 2, 3, 4}, 2))
	//fmt.Println(bs([]int{1, 2, 3, 4}, 1))
	//fmt.Println(bs([]int{1, 2, 3, 4}, 0))
}
