package problem0372

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a   int
	b   []int
	ans int
}{

	// 可以有多个 testcase
}

func Test_superPow(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, superPow(tc.a, tc.b), "输入:%v", tc)
	}
}

func Benchmark_superPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			superPow(tc.a, tc.b)
		}
	}
}
