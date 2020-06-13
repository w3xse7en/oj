package problem0018

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums   []int
	target int
	ans    [][]int
}{
	{[]int{1, 2, 2, 3, 0, 0}, 4, [][]int{{1, 3, 0, 0}, {2, 2, 0, 0}}},
	//{[]int{0,0,0,0,0,0,1},0,[][]int{{0,0,0,0}}},

	// 可以有多个 testcase
}

type A struct {
	a, b, c, d int
}

func Test_eq(t *testing.T) {

	fmt.Println(A{
		a: 1,
		b: 0,
		c: 0,
		d: 0,
	} == A{
		a: 1,
		b: 0,
		c: 0,
		d: 0,
	})
}
func Test_fourSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, fourSum(tc.nums, tc.target), "输入:%v", tc)
	}
}

func Benchmark_fourSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			fourSum(tc.nums, tc.target)
		}
	}
}
