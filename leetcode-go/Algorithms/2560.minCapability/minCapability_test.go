package _560_minCapability

import "testing"

func Test_minCapability(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 3, 5, 9},
				k:    2,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				nums: []int{7, 3, 9, 5},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCapability(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minCapability() = %v, want %v", got, tt.want)
			}
		})
	}
}
