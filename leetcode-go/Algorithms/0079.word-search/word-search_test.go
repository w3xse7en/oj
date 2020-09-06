package problem0079

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var b = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'C', 'S'},
	{'A', 'D', 'E', 'E'},
}
var c = [][]byte{
	{'a', 'b'},
	{'c', 'd'},
}
var d = [][]byte{
	{'A', 'B', 'C', 'E'},
	{'S', 'F', 'E', 'S'},
	{'A', 'D', 'E', 'E'},
}

// tcs is testcase slice
var tcs = []struct {
	board [][]byte
	word  string
	ans   bool
}{
	{d, "ABCEFSADEESE", true},
	{b, "ABCCED", true},
	{b, "SEE", true},
	{b, "ABCB", false},
	{c, "abcd", false},
	// 可以有多个 testcase
}

func Test_exist(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, exist(tc.board, tc.word), "输入:%v", tc)
	}
}

func Benchmark_exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			exist(tc.board, tc.word)
		}
	}
}
