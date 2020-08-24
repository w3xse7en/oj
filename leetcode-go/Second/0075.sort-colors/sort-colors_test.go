package problem0075

import (
	"fmt"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
}{
	{[]int{2, 0, 1}},
	{[]int{2, 1, 2, 0, 1, 2, 1, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0, 2, 1, 0}},

	// 可以有多个 testcase
}

func Test_sortColors(t *testing.T) {
	for _, tc := range tcs {
		sortColors(tc.nums)
		fmt.Println(tc)
	}
}

func Benchmark_sortColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortColors(tc.nums)
		}
	}
}
