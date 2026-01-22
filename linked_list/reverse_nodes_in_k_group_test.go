package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		{
			name: "nil head",
			head: nil,
			k:    3,
			want: nil,
		},
		{
			name: "k equals 1",
			head: []int{1, 2, 3, 4},
			k:    1,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "k equals length",
			head: []int{1, 2, 3},
			k:    3,
			want: []int{3, 2, 1},
		},
		{
			name: "multiple full groups",
			head: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: []int{3, 2, 1, 6, 5, 4},
		},
		{
			name: "last group smaller",
			head: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},
		{
			name: "k greater than length",
			head: []int{1, 2},
			k:    3,
			want: []int{1, 2},
		},
		{
			name: "duplicates",
			head: []int{1, 1, 2, 2, 3, 3},
			k:    2,
			want: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.head)
			gotHead := reverseKGroup(head, tt.k)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
