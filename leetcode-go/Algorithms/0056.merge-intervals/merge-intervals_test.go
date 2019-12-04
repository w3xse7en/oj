package problem0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	intervals [][]int
	ans       [][]int
}{
	{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
	{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
	{[][]int{{1, 4}, {0, 1}}, [][]int{{0, 4}}},
	{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	{[][]int{{2, 3}, {1, 4}}, [][]int{{1, 4}}},
	{[][]int{{1, 4}, {0, 0}}, [][]int{{0, 0}, {1, 4}}},
	{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}}},
	{[][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}, [][]int{{2, 4}, {5, 5}}},

	// 可以有多个 testcase
}

func Test_merge(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, merge(tc.intervals), "输入:%v", tc)
	}
}

func Benchmark_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			merge(tc.intervals)
		}
	}
}
