package problem0060

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	k   int
	ans string
}{
	{1, 1, "1"},
	//{2, 3, ""},
	{3, 3, "213"},
	{4, 9, "2314"},
	//{5, 3, ""},
	//{6, 3, ""},
	//{7, 3, ""},
	//{8, 3, ""},
	{9, 219601, "647123589"},

	// 可以有多个 testcase
}

func Test_getPermutation(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, getPermutation(tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_getPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getPermutation(tc.n, tc.k)
		}
	}
}
