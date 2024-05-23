package problem0167

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_idx(t *testing.T) {
	type args struct {
		nums []int
		t    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
				t:    4,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				t:    2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
				t:    3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, idx(tt.args.nums, tt.args.t), "idx(%v, %v)", tt.args.nums, tt.args.t)
		})
	}
}
