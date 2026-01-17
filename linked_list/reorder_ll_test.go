package linked_list

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
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
			want: []int{1, 2},
		},
		{
			name: "three nodes",
			in:   []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "even length example",
			in:   []int{2, 4, 6, 8},
			want: []int{2, 8, 4, 6},
		},
		{
			name: "odd length example",
			in:   []int{2, 4, 6, 8, 10},
			want: []int{2, 10, 4, 8, 6},
		},
		{
			name: "repeated values",
			in:   []int{5, 5, 5, 5},
			want: []int{5, 5, 5, 5},
		},
		{
			name: "longer list",
			in:   []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 6, 2, 5, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.in)
			reorderList(head)
			got := sliceFromList(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
