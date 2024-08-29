package _153_findMin

import "testing"

func Test_findMin(t *testing.T) {
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
			args: args{[]int{1, 2, 3, 4, 5, 6, 7}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{4, 5, 6, 7, 1, 2, 3}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{7, 1, 2, 3, 4, 5, 6}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{7, 1}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
