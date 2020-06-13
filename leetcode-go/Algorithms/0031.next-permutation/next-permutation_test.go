package problem0031

import (
	"fmt"
	"testing"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
}{
	{[]int{23, 16, 1, 20}},
	{[]int{1, 2}},
	{[]int{2, 3, 1}},
	{[]int{1, 3, 3, 3, 2, 2, 2, 2, 3, 3}},
	{[]int{1, 3, 3, 3, 3}},
	{[]int{3, 3, 3, 3, 2, 1}},
	{[]int{3, 4, 4, 4, 4, 4}},
	{[]int{1, 9, 8, 7, 6, 5, 4}},
	{[]int{1, 3, 2}},
	{[]int{5, 4, 7, 5, 3, 2}},
	// 可以有多个 testcase
}

func Test_nextPermutation(t *testing.T) {
	for _, tc := range tcs {
		nextPermutation(tc.nums)
		fmt.Println(tc.nums)
	}
}
func Benchmark_nextPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			nextPermutation(tc.nums)
		}
	}
}
