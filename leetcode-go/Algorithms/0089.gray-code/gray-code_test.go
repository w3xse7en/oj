package problem0089

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans []int
}{
	//{2, []int{0,2,3,1}},
	{3, []int{0, 1, 3, 2, 6, 7, 5, 4}},

	// 可以有多个 testcase
}

func Test_b(t *testing.T) {
	fmt.Println(1 << 1)
	fmt.Println(1 << 2)
	fmt.Println(1 << 3)
}
func Test_grayCode(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, grayCode(tc.n), "输入:%v", tc)
	}
}

func Benchmark_grayCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			grayCode(tc.n)
		}
	}
}
