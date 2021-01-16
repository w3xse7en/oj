package problem0567

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s1  string
	s2  string
	ans bool
}{
	{"a", "a", true},
	{"aa", "a", false},
	{"aa", "c", false},
	{"a", "aa", true},
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},

	// 可以有多个 testcase
}

func Test_checkInclusion(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, checkInclusion(tc.s1, tc.s2), "输入:%v", tc)
	}
}

func Benchmark_checkInclusion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			checkInclusion(tc.s1, tc.s2)
		}
	}
}
