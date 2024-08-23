package _394_decodeString

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{s: "100[leetcode]"},
			want: "",
		},
		{
			name: "",
			args: args{s: "3[a]2[bc]"},
			want: "",
		},
		{
			name: "",
			args: args{s: "3[a2[bc]]"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
