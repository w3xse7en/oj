package problem0130

import (
	"fmt"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
}{
	{[][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}},
	// 可以有多个 testcase
}

func Test_solve(t *testing.T) {
	for _, tc := range tcs {
		solve(tc.board)
		fmt.Println(tc.board)
	}
}

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			solve(tc.board)
		}
	}
}
