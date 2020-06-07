package problem0016

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    int
}{
	{[]int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33}, 0, 0},
	{[]int{1, 643653, 7756, 765, 123, 542, 53456, 7457, 8467, 68, 566, 34, 534, 5}, 10, 40},
	{[]int{-1, 2, 1, -4}, 1, 2},
	{[]int{-1, 2, 1}, 1, 2},
	// 可以有多个 testcase
}

func Test_one(t *testing.T) {
	ints := []int{1, 643653, 7756, 765, 123, 542, 53456, 7457, 8467, 68, 566, 34, 534, 5, 3}
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(oneClosest(ints, 4))
}

func Test_threeSumClosest(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, threeSumClosest(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_threeSumClosest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			threeSumClosest(tc.nums, tc.target)
		}
	}
}
