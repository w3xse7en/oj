package problem0093

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans []string
}{
	{"666666254", []string{"66.66.66.254"}},
	{"0255255254", []string{"0.255.255.254"}},
	//{"25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	//{"0000", []string{"0.0.0.0"}},
	//{"010010", []string{"0.10.0.10","0.100.1.0"}},
	//{"101023", []string{"1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"}},

	// 可以有多个 testcase
}

func Test_restoreIpAddresses(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, restoreIpAddresses(tc.s), "输入:%v", tc)
	}
}

func Benchmark_restoreIpAddresses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			restoreIpAddresses(tc.s)
		}
	}
}
