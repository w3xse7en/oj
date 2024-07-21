package _400_findNthDigit

import "testing"

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{10},
			want: 1,
		},
		{
			name: "",
			args: args{1},
			want: 1,
		},
		{
			name: "",
			args: args{11},
			want: 0,
		},
		{
			name: "",
			args: args{100},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigit(tt.args.n); got != tt.want {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
