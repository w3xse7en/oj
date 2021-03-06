package problem0043

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	num1 string
	num2 string
	ans  string
}{
	{"123", "456", "56088"},
	// {"1","999","999"},
	// {"1","0","0"},
	{"0", "510", "0"},
	// {"99999","99999","9999800001"},

	// 可以有多个 testcase
}

func Test_multiply(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, multiply(tc.num1, tc.num2), "输入:%v", tc)
	}
}

func Benchmark_multiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			multiply(tc.num1, tc.num2)
		}
	}
}
