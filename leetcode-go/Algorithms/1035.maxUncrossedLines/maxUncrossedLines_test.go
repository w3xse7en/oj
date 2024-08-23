package _035_maxUncrossedLines

import "testing"

func Test_maxUncrossedLines(t *testing.T) {
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
			args: args{
				nums1: []int{3, 1, 4, 1, 1, 3, 5, 1, 2, 2},
				nums2: []int{4, 1, 5, 2, 1, 1, 1, 5, 3, 1, 1, 1, 2, 3, 1, 4, 3, 5, 5, 3, 1, 2, 3, 2, 4, 1, 1, 1, 5, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxUncrossedLines(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("maxUncrossedLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
