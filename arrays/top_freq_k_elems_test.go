package arrays

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "example 1: k=2 from mixed frequencies",
			nums: []int{1, 2, 2, 3, 3, 3},
			k:    2,
			want: []int{2, 3},
		},
		{
			name: "example 2: single element repeated",
			nums: []int{7, 7},
			k:    1,
			want: []int{7},
		},
		{
			name: "k=1 from multiple elements",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    1,
			want: []int{1},
		},
		//{
		//	name: "all elements same frequency",
		//	nums: []int{1, 2, 3, 4},
		//	k:    2,
		//	want: []int{1, 2}, // any 2 elements are valid
		//},
		{
			name: "k equals distinct elements",
			nums: []int{1, 1, 2, 2, 3, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "single element array",
			nums: []int{5},
			k:    1,
			want: []int{5},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, -2, -2, -3},
			k:    2,
			want: []int{-1, -2},
		},
		{
			name: "mixed positive and negative",
			nums: []int{4, 1, -1, 2, -1, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
		{
			name: "large frequency difference",
			nums: []int{1, 1, 1, 1, 1, 2, 3},
			k:    1,
			want: []int{1},
		},
		{
			name: "zeros included",
			nums: []int{0, 0, 0, 1, 1, 2},
			k:    2,
			want: []int{0, 1},
		},
		{
			name: "k equals 1 with many elements",
			nums: []int{5, 3, 1, 1, 1, 3, 5, 73, 1},
			k:    1,
			want: []int{1},
		},
		{
			name: "multiple elements with descending frequencies",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if len(got) != tt.k {
				t.Errorf("topKFrequent() returned %d elements, want %d", len(got), tt.k)
				return
			}
			if !equalIgnoreOrder(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalIgnoreOrder compares two slices for equality, ignoring order
func equalIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)

	sort.Ints(aCopy)
	sort.Ints(bCopy)

	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}

	return true
}
