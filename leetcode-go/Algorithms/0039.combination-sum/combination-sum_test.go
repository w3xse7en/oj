package problem0039

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	candidates []int
	target     int
	ans        [][]int
}{
	//{[]int{2, 3, 6, 7}, 7, [][]int{}},
	{[]int{2, 3, 5}, 8, [][]int{}},
	{[]int{7, 3, 2}, 18, [][]int{}},

	// 可以有多个 testcase
}

func Test_combinationSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, combinationSum(tc.candidates, tc.target), "输入:%v", tc)
	}
}

func Benchmark_combinationSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			combinationSum(tc.candidates, tc.target)
		}
	}
}

func Test_append(t *testing.T) {
	f := func() []int {
		return nil
	}
	fmt.Println(append(f(), 3))
}
