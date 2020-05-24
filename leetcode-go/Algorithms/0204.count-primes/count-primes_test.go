package problem0204

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
	{10, 4},
	{2, 0},
	{3, 1},
	{65, 18},
	{106, 27},
	// 可以有多个 testcase
}

func Test_countPrimel(t *testing.T) {
	countPrimes(1000)
}
func Test_countPrime(t *testing.T) {
	s := map[int]int{}
	for i := 3; i < 1000; i++ {
		s[countPrimes(i)]++
	}
	for i := 0; i <= 200; i++ {
		fmt.Println(i, s[i])
	}
}
func Test_countPrimes(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, countPrimes(tc.n), "输入:%v", tc)
	}
}

func Benchmark_countPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countPrimes(tc.n)
		}
	}
}
