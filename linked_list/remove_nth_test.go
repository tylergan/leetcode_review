package linked_list

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		n    int
		want []int
	}{
		{
			name: "single node removes head",
			in:   []int{5},
			n:    1,
			want: nil,
		},
		{
			name: "remove tail",
			in:   []int{1, 2, 3},
			n:    1,
			want: []int{1, 2},
		},
		{
			name: "remove head",
			in:   []int{1, 2, 3},
			n:    3,
			want: []int{2, 3},
		},
		{
			name: "remove middle",
			in:   []int{1, 2, 3, 4},
			n:    2,
			want: []int{1, 2, 4},
		},
		{
			name: "two nodes remove first",
			in:   []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "two nodes remove last",
			in:   []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "repeated values",
			in:   []int{7, 7, 7, 7},
			n:    3,
			want: []int{7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listFromSlice(tt.in)
			gotHead := removeNthFromEnd(head, tt.n)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
