package _002_maximumSetSize

import "testing"

func Test_maximumSetSize(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[]int{1, 2, 1, 2}, []int{1, 1, 1, 1}},
			want: 2,
		},
		{
			name: "",
			args: args{[]int{1, 2, 3, 4, 5, 6}, []int{2, 3, 2, 3, 2, 3}},
			want: 5,
		},
		{
			name: "",
			args: args{[]int{1, 1, 2, 2, 3, 3}, []int{4, 4, 5, 5, 6, 6}},
			want: 6,
		},
		{
			name: "",
			args: args{[]int{12, 23, 41, 9}, []int{1, 1, 1, 1}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSetSize(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maximumSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
