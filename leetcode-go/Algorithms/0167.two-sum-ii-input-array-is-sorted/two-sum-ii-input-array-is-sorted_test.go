package problem0167

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numbers []int
	target  int
	ans     []int
}{
	{numbers: []int{0, 0, 3, 4}, target: 0, ans: []int{1, 2}},
	{numbers: []int{2, 7, 11, 15}, target: 9, ans: []int{1, 2}},
	{numbers: []int{1, 2, 3, 17, 18, 24, 26, 38, 47, 49, 55, 70}, target: 117, ans: []int{9, 12}},

	// 可以有多个 testcase
}

func Test_twoSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, twoSum(tc.numbers, tc.target), "输入:%v", tc)
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			twoSum(tc.numbers, tc.target)
		}
	}
}
