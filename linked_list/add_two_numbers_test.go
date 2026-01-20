package linked_list

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "single digit no carry",
			l1:   []int{3},
			l2:   []int{4},
			want: []int{7},
		},
		{
			name: "single digit with carry",
			l1:   []int{9},
			l2:   []int{9},
			want: []int{8, 1},
		},
		{
			name: "same length no carry",
			l1:   []int{1, 2, 3},
			l2:   []int{4, 5, 6},
			want: []int{5, 7, 9},
		},
		{
			name: "same length with carry chain",
			l1:   []int{9, 9, 9},
			l2:   []int{1, 0, 0},
			want: []int{0, 0, 0, 1},
		},
		{
			name: "different lengths no final carry",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6},
			want: []int{7, 0, 4},
		},
		{
			name: "different lengths with final carry",
			l1:   []int{5, 6},
			l2:   []int{5, 4, 9},
			want: []int{0, 1, 0, 1},
		},
		{
			name: "zeros",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := listFromSlice(tt.l1)
			l2 := listFromSlice(tt.l2)
			gotHead := addTwoNumbers(l1, l2)
			got := sliceFromList(gotHead)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
