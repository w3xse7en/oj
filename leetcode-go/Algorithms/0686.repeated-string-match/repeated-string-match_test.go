package problem0686

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   string
	B   string
	ans int
}{

	// 可以有多个 testcase
}

func Test_repeatedStringMatch(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, repeatedStringMatch(tc.A, tc.B), "输入:%v", tc)
	}
}

func Benchmark_repeatedStringMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			repeatedStringMatch(tc.A, tc.B)
		}
	}
}
