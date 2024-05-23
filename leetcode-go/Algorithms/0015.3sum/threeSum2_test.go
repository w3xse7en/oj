package problem0015

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_half(t *testing.T) {
	type args struct {
		nums []int
		t    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{[]int{1, 2, 3, 4}, 3},
			want: true,
		},
		{
			name: "",
			args: args{[]int{0}, 0},
			want: true,
		},
		{
			name: "",
			args: args{[]int{1, 2}, 0},
			want: false,
		},
		{
			name: "",
			args: args{[]int{1, 2, 4, 5}, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, half(tt.args.nums, tt.args.t), "half(%v, %v)", tt.args.nums, tt.args.t)
		})
	}
}
