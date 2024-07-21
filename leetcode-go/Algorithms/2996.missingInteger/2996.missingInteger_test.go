package _996_missingInteger

import "testing"

func Test_missingInteger(t *testing.T) {
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
			args: args{[]int{3, 4, 5, 1, 12, 14, 13}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingInteger(tt.args.nums); got != tt.want {
				t.Errorf("missingInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
