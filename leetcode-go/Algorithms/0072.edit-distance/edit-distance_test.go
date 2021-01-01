package problem0072

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	word1 string
	word2 string
	ans   int
}{

	{"abc", "abc", 0},
	{"abc", "cba", 2},
	{"horse", "ros", 3},
	{"intention", "execution", 5},
	{"tioninten", "execution", 8},
	{"tionn", "executisdg", 8},

	// 可以有多个 testcase
}

func Test_minDistance(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minDistance(tc.word1, tc.word2), "输入:%v", tc)
	}
}

func Benchmark_minDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDistance(tc.word1, tc.word2)
		}
	}
}
