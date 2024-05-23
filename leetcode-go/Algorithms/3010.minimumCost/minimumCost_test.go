package _010_minimumCost

import "testing"

func Test_minimumCost(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[]int{1, 2, 1, 1}},
			want: 3,
		},
		{
			name: "",
			args: args{[]int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "",
			args: args{[]int{9, 2, 3, 1}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumCost(tt.args.nums); got != tt.want {
				t.Errorf("minimumCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
