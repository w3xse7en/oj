package problem0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	k    int
	ans  []int
}{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{-1}, 2, []int{-1}},
	//{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	//{[]int{1, 2, 3, 4, 5, 6, 7}, 4, []int{5, 6, 7, 1, 2, 3, 4}},
	// 可以有多个 testcase
}

func Test_rotate(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		rotate(tc.nums, tc.k)
		a.Equal(tc.ans, tc.nums, "输入:%v", tc)
	}
}

func Benchmark_rotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			rotate(tc.nums, tc.k)
		}
	}
}
