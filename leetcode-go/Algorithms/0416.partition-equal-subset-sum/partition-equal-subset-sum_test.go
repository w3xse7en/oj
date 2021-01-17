package problem0416

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{
	{[]int{1, 5, 11, 5}, true},
	{[]int{1, 2, 3, 5}, false},
	{[]int{1, 2, 4, 5}, true},
	{[]int{1, 2}, false},
	{[]int{1}, false},

	// 可以有多个 testcase
}

func Test_canPartition(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, canPartition(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_canPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			canPartition(tc.nums)
		}
	}
}
