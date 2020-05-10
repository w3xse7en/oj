package problem0172

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{
	//{7, 1},
	//{3, 0},
	//{5, 1},
	//{10, 2},
	//{30, 7},
	{532, 7},

	// 可以有多个 testcase
}

func Test_a(t *testing.T) {
	abc(5)
	abc(25)
	abc(24)
	abc(50)
	abc(100)
	abc(125)
	abc(250)
	abc(443)
	abc(482)
	abc(499)
}
func abc(n int) int {
	fmt.Println(n, fmt.Sprintf("%v/%v=%v", n, 5, n/5), fmt.Sprintf("%v/%v=%v", n, 25, n/25), fmt.Sprintf("%v/%v=%v", n, 125, n/125), n/5+n/25+n/125)
	r, step := 0, 1
	for n/step > 0 {
		step *= 5
		r += n / step
	}
	return r
}
func Test_m(t *testing.T) {
	b, a := []int{5}, []int{2}
	r := multiple(a, b)
	fmt.Println(r)
}

func Test_itob(t *testing.T) {
	n := 123
	r := iToArray(n)
	fmt.Println(r)
}

func Test_r(t *testing.T) {
	mp := map[int]int{}
	for i := 0; i < 500; i++ {
		zeroes := trailingZeroes(i)
		fmt.Println(i, zeroes)
		mp[zeroes]++
	}
	s := make([]int, 0)
	for k := range mp {
		s = append(s, k)
	}
	sort.Ints(s)

	for _, k := range s {
		fmt.Println(k, mp[k])
	}
}

func Test_trailingZeroes(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, trailingZeroes(tc.n), "输入:%v", tc)
	}
}

func Benchmark_trailingZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trailingZeroes(tc.n)
		}
	}
}
