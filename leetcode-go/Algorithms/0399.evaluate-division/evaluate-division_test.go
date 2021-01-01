package problem0399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	equations [][]string
	values    []float64
	queries   [][]string
	ans       []float64
}{
	//{
	//	equations: [][]string{{"a", "b"}, {"b", "c"}},
	//	values:    []float64{2.0, 3.0},
	//	queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
	//	ans:       []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
	//},
	//{
	//	equations: [][]string{{"a", "e"}, {"b", "e"}},
	//	values:    []float64{4.0, 3.0},
	//	queries:   [][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}},
	//	ans:       []float64{1.33333,1.00000,-1.00000},
	//},
	{
		equations: [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}},
		values:    []float64{3.0, 4.0, 5.0, 6.0},
		queries:   [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}},
		ans:       []float64{360.00000, 0.00833, 20.00000, 1.00000, -1.00000, -1.00000},
	},
	// 可以有多个 testcase
}

func Test_calcEquation(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, calcEquation(tc.equations, tc.values, tc.queries), "输入:%v", tc)
	}
}

func Benchmark_calcEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			calcEquation(tc.equations, tc.values, tc.queries)
		}
	}
}
