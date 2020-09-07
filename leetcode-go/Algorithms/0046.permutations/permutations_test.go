package problem0046

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  [][]int
}{
	//{[]int{1,2,3}, [][]int{}},
	//{[]int{3,2,1}, [][]int{}},
	{[]int{1, 2, 3, 4}, [][]int{}},
	// 可以有多个 testcase
}

func Test_s(t *testing.T) {
	ints := []int{1, 3, 5, 7, 8}
	var a []int
	for i := range ints {
		if ints[i] == 9 {
			a = append(ints[:i], ints[i+1:]...)
		}
	}
	fmt.Println(a)
}
func Test_permute(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, permute(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_permute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			permute(tc.nums)
		}
	}
}
