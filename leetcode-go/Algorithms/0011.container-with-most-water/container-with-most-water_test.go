package problem0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	height []int
	ans    int
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 16},
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 16},
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{7, 3, 8, 4, 5, 2, 6, 8, 1}, 49},

	// 可以有多个 testcase
}

func Test_maxArea(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxArea(tc.height), "输入:%v", tc)
	}
}

func Benchmark_maxArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxArea(tc.height)
		}
	}
}
