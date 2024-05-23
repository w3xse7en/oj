package _055_canJump

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{[]int{2, 3, 1, 1, 4}},
			want: true,
		}, {
			name: "",
			args: args{[]int{3, 2, 1, 0, 4}},
			want: false,
		}, {
			name: "",
			args: args{[]int{1, 3, 0, 0, 0, 0}},
			want: false,
		}, {
			name: "",
			args: args{[]int{1, 3, 0, 0, 1, 0}},
			want: true,
		}, {
			name: "",
			args: args{[]int{0, 1}},
			want: false,
		}, {
			name: "",
			args: args{[]int{0}},
			want: true,
		}, {
			name: "",
			args: args{[]int{1, 1, 0, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
