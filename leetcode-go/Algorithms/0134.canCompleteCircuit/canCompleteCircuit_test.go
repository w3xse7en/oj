package _134_canCompleteCircuit

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			want: 3,
		},
		{
			name: "",
			args: args{[]int{2, 3, 4}, []int{3, 4, 3}},
			want: -1,
		},
		{
			name: "",
			args: args{[]int{2, 3, 4}, []int{2, 3, 4}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{1}, []int{1}},
			want: 0,
		},
		{
			name: "",
			args: args{[]int{1, 2}, []int{1, 3}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
