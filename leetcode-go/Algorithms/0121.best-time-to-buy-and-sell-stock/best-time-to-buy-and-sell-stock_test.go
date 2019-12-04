package problem0121

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	prices []int
	ans    int
}{
	{[]int{5, 4, 3, 2, 1}, 0},
	{[]int{7, 1, 5, 3, 6, 4}, 5},

	// 可以有多个 testcase
}

func Test_maxProfit(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maxProfit(tc.prices), "输入:%v", tc)
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxProfit(tc.prices)
		}
	}
}
