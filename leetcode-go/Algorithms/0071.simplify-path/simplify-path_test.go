package problem0071

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	path string
	ans  string
}{
	{"/a//b////c/d//././/..", "/a/b/c"},
	{"/../", "/"},
	{"/home//foo/", "/home/foo"},
	{"/a/./b/../../c/", "/c"},
	{"/a/../../b/../c//.//", "/c"},

	// 可以有多个 testcase
}

func Test_simplifyPath(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, simplifyPath(tc.path), "输入:%v", tc)
	}
}

func Benchmark_simplifyPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			simplifyPath(tc.path)
		}
	}
}
