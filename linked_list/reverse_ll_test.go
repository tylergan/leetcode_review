package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "empty list",
			in:   nil,
			want: nil,
		},
		{
			name: "single node",
			in:   []int{7},
			want: []int{7},
		},
		{
			name: "two nodes",
			in:   []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "multiple nodes",
			in:   []int{0, 1, 2, 3},
			want: []int{3, 2, 1, 0},
		},
		{
			name: "negative values",
			in:   []int{-1, -2, -3},
			want: []int{-3, -2, -1},
		},
		{
			name: "repeated values",
			in:   []int{5, 5, 5},
			want: []int{5, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.in)
			gotHead := reverseList(head)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func listFromSlice(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	curr := head
	for _, v := range values[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func sliceFromList(head *ListNode) []int {
	var out []int
	for curr := head; curr != nil; curr = curr.Next {
		out = append(out, curr.Val)
	}
	return out
}
