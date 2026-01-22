package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{
			name:  "nil lists",
			lists: nil,
			want:  nil,
		},
		{
			name:  "empty lists slice",
			lists: [][]int{},
			want:  nil,
		},
		{
			name:  "single empty list",
			lists: [][]int{{}},
			want:  nil,
		},
		{
			name:  "single non-empty list",
			lists: [][]int{{1, 2, 3}},
			want:  []int{1, 2, 3},
		},
		{
			name:  "example interleaved",
			lists: [][]int{{1, 2, 4}, {1, 3, 5}, {3, 6}},
			want:  []int{1, 1, 2, 3, 3, 4, 5, 6},
		},
		{
			name:  "with nil and empty lists",
			lists: [][]int{{}, {0}, nil, {2, 2}},
			want:  []int{0, 2, 2},
		},
		{
			name:  "negatives and duplicates",
			lists: [][]int{{-5, -3, -3, 1}, {-4, -3, 2}, {-2, 0, 0}},
			want:  []int{-5, -4, -3, -3, -3, -2, 0, 0, 1, 2},
		},
		{
			name:  "many singletons",
			lists: [][]int{{5}, {1}, {3}, {2}, {4}},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var listNodes []*ListNode
			if tt.lists != nil {
				listNodes = make([]*ListNode, 0, len(tt.lists))
				for _, list := range tt.lists {
					listNodes = append(listNodes, listFromSlice(list))
				}
			}

			gotHead := mergeKLists(listNodes)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
