package _721_swapNodes

import (
	"reflect"
	"testing"
)

func Test_swapNodes(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: BuildList([]int{1, 2, 3, 4, 5}),
				k:    2,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: BuildList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}),
				k:    5,
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				head: BuildList([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5}),
				k:    1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapNodes(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildList(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	head := &ListNode{Val: is[0]}
	node := head
	for k, v := range is {
		if k > 0 {
			node.Next = &ListNode{Val: v}
			node = node.Next
		}
	}
	return head
}
