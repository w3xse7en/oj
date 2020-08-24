package problem0074

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = [][]int{
	{},
}

// tcs is testcase slice
var tcs = []struct {
	matrix [][]int
	target int
	ans    bool
}{
	{m, 1, false},
	//{m, 2, false},
	//{m, 0, false},
	//{m, 10, true},
	//{m, 20, true},
	//{m, 13, false},
	//{m, 9, false},
	//{m, 21, false},
	//{m, 100, false},
	//{m, 77, false},
	//{m, 34, true},
	//{m, 50, true},
	// 可以有多个 testcase
}

func Test_searchRow(t *testing.T) {
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 0))
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 2))
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 4))
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 7))
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 44))
	fmt.Println(searchRow([]int{1, 3, 5, 6, 10, 45, 66, 77, 88, 99, 101}, 102))
}
func Test_searchMatrix(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, searchMatrix(tc.matrix, tc.target), "输入:%v", tc)
	}
}

func Benchmark_searchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			searchMatrix(tc.matrix, tc.target)
		}
	}
}
