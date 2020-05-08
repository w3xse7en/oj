package problem0168

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans string
}{

	{731, "ABC"},
	{1, "A"},
	{28, "AB"},
	{701, "ZY"},
	{18277, "ZZY"},
	{1377, "AZY"},

	// 可以有多个 testcase
}

func Test_c(t *testing.T) {
	fmt.Println(string(byte(65)))
	fmt.Println("A", 1/27, 1%27)
	fmt.Println("Z", 26/26%26, 26%26)
	fmt.Println("AA", 27%26, 27/26%26)
	fmt.Println("AB", 28%26, 28/26%26)
	fmt.Println("AC", 29%26, 29/26%26)
	fmt.Println("ZY", 701/26%26, 701%26)
	fmt.Println("ZZ", 702/26/26/26%26, 702/26/26%26, 702/26%26, 702%26)
	fmt.Println("AAA", 703/27, 703%27)
	fmt.Println("ABC", 731/27, 731/27%27, 731%27)
	fmt.Println("ABC", 731/26/26/26-3%26, 731/26/26-2%26, 731/26-1%26, 731%26)
	fmt.Println("AZY", 1*26*26+26*26+25)
	fmt.Println("AZY", 1377/26/26/26%26, 1377/26/26%26, 1377/26%26, 1377%26)
}
func Test_convertToTitle(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, convertToTitle(tc.n), "输入:%v", tc)
	}
}

func Benchmark_convertToTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			convertToTitle(tc.n)
		}
	}
}
