package problem0073

import (
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
}{
	{},
	// 可以有多个 testcase
}

func Test_setZeroes(t *testing.T) {
	for _, tc := range tcs {
		setZeroes(tc.matrix)
	}
}

func Benchmark_setZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			setZeroes(tc.matrix)
		}
	}
}
