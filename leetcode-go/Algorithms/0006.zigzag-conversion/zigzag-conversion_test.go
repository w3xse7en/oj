package problem0006

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s       string
	numRows int
	ans     string
}{

	{"PAYPALISHIRING", 2, "PYAIHRNAPLSIIG"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 0, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"P123", 4, "P123"},

	// 可以有多个 testcase
}

func Test_convert(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, convert(tc.s, tc.numRows), "输入:%v", tc)
	}
}

func Benchmark_convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			convert(tc.s, tc.numRows)
		}
	}
}
