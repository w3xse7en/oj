package _068_fullJustify

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				words:    []string{"Listen", "to", "many,", "speak", "to", "a", "few."},
				maxWidth: 6,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
