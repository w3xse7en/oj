package problem0003

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans int
}{
	//{"dvdf", 3},
	//{"aab", 2},
	//{"bbbbb", 1},
	//{"abcabcbb", 3},
	//{"pwwkew", 3},
	{"aaabbbcccdddeeeffgh", 3},

	// 可以有多个 testcase
}

func Test_index(t *testing.T) {
	fmt.Println(strings.Index("", "d"))
}
func Test_lengthOfLongestSubstring(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, lengthOfLongestSubstring(tc.s), "输入:%v", tc)
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lengthOfLongestSubstring(tc.s)
		}
	}
}
