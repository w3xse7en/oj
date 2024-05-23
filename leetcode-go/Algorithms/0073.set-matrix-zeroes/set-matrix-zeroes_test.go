package problem0073

import (
	"fmt"
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

func Test_setZeroes1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{[][]int{{1, 0, 3}, {2, 1, 4}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			fmt.Println(tt.args.matrix)
		})
	}
}
