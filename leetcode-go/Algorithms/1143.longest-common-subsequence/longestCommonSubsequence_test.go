package problem1143

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonSubsequence1(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				text1: "bsbininm",
				text2: "jmjkbkjkv",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestCommonSubsequence(tt.args.text1, tt.args.text2), "longestCommonSubsequence(%v, %v)", tt.args.text1, tt.args.text2)
		})
	}
}
