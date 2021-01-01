package problem0390

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{
	//{1000000, 2},
	//{2, 2},
	//{3, 2},
	//{4, 2},
	{9, 6},
	//{1, 1},
	// 可以有多个 testcase
}

func Test_f(t *testing.T) {
	for i := 1; i < 1000; i++ {
		fmt.Println(i, lastRemaining(i))
	}
}
func Test_fo(t *testing.T) {

}
func Test_pre(t *testing.T) {
}
func Test_lastRemaining(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, lastRemaining(tc.n), "输入:%v", tc)
	}
}

func Benchmark_lastRemaining(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lastRemaining(tc.n)
		}
	}
}
