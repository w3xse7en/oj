package problem1143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	text1 string
	text2 string
	ans   int
}{

	//{"abcde", "ace", 3},
	{"bsbininm", "jmjkbkjkv", 1},
	// 可以有多个 testcase
}

func Test_longestCommonSubsequence(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, longestCommonSubsequence(tc.text1, tc.text2), "输入:%v", tc)
	}
}

func Benchmark_longestCommonSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			longestCommonSubsequence(tc.text1, tc.text2)
		}
	}
}
