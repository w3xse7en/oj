package _023_mergeKLists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{[]*ListNode{{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			}, {
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			}, {
				Val: 2,
				Next: &ListNode{
					Val:  6,
					Next: nil,
				},
			},
			}},
			want: nil,
		}, {
			name: "",
			args: args{[]*ListNode{{
				Val:  2,
				Next: nil,
			}, nil, {
				Val:  -1,
				Next: nil,
			}, nil,
			}},
			want: nil,
		}, {
			name: "",
			args: args{[]*ListNode{nil}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestName(t *testing.T) {
	list := []int{1, 2, 3}
	fmt.Println(list[1:])
}
