package _209_minSubArrayLen

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				target: 10,
				nums:   []int{11, 5, 5},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				target: 10,
				nums:   []int{1, 2, 3, 4, 5, 5, 5},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				target: 11,
				nums:   []int{1, 2, 3, 4, 5, 5, 5},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
