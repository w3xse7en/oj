package problem0735

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	asteroids []int
	ans       []int
}{
	{asteroids: []int{-2, -1, 1, 2}, ans: []int{-2, -1, 1, 2}},
	{asteroids: []int{10, 2, -5}, ans: []int{10}},
	{asteroids: []int{1, -2, -2, -2}, ans: []int{-2, -2, -2}},
	{asteroids: []int{5, 10, -5}, ans: []int{5, 10}},
	{asteroids: []int{8, -8}, ans: []int{}},
	// 可以有多个 testcase
}

func Test_f(t *testing.T) {
	fmt.Println(f([]int{-12, -3}, -3))
	fmt.Println(f([]int{12, 3}, -3))
	fmt.Println(f([]int{-12, 3}, -3))
	fmt.Println(f([]int{12, 3}, -9))
	fmt.Println(f([]int{12, 3}, -12))
	fmt.Println(f([]int{12}, -12))
}
func Test_asteroidCollision(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, asteroidCollision(tc.asteroids), "输入:%v", tc)
	}
}

func Benchmark_asteroidCollision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			asteroidCollision(tc.asteroids)
		}
	}
}
