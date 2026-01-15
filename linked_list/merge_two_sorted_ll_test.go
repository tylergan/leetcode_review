package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{
			name:  "both empty",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "list1 empty",
			list1: nil,
			list2: []int{1, 2},
			want:  []int{1, 2},
		},
		{
			name:  "list2 empty",
			list1: []int{3, 4},
			list2: nil,
			want:  []int{3, 4},
		},
		{
			name:  "interleaved",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 5},
			want:  []int{1, 1, 2, 3, 4, 5},
		},
		{
			name:  "list1 all smaller",
			list1: []int{-3, -2, -1},
			list2: []int{0, 1},
			want:  []int{-3, -2, -1, 0, 1},
		},
		{
			name:  "list2 all smaller",
			list1: []int{4, 5},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "duplicates across lists",
			list1: []int{1, 2, 2, 4},
			list2: []int{2, 2, 3},
			want:  []int{1, 2, 2, 2, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := listFromSlice(tt.list1)
			list2 := listFromSlice(tt.list2)
			gotHead := mergeTwoLists(list1, list2)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
