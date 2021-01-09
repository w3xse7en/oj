package problem0221

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]byte
	ans    int
}{
	{matrix: [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}, ans: 4},
	{matrix: [][]byte{
		{'1', '0'},
		{'1', '0'},
	}, ans: 1},
	{matrix: [][]byte{
		{'0', '0'},
	}, ans: 0},
	{matrix: [][]byte{
		{'0'},
	}, ans: 0},

	// 可以有多个 testcase
}

func Test_maximalSquare(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, maximalSquare(tc.matrix), "输入:%v", tc)
	}
}

func Benchmark_maximalSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maximalSquare(tc.matrix)
		}
	}
}
