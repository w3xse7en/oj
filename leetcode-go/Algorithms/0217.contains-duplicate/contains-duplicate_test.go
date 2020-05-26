package problem0217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{
	{[]int{1, 1}, true},
	{[]int{0, 1}, false},
	{[]int{1, 1, 2, 2}, true},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},

	// 可以有多个 testcase
}

func Test_containsDuplicate(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, containsDuplicate(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_containsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containsDuplicate(tc.nums)
		}
	}
}
