package backtrack

import (
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "example 1 - multiple combinations",
			nums:   []int{2, 5, 6, 9},
			target: 9,
			want:   [][]int{{2, 2, 5}, {9}},
		},
		{
			name:   "example 2 - repeated use of elements",
			nums:   []int{3, 4, 5},
			target: 16,
			want:   [][]int{{3, 3, 3, 3, 4}, {3, 3, 5, 5}, {3, 4, 4, 5}, {4, 4, 4, 4}},
		},
		{
			name:   "example 3 - no valid combination",
			nums:   []int{3},
			target: 5,
			want:   [][]int{},
		},
		{
			name:   "single element equals target",
			nums:   []int{7},
			target: 7,
			want:   [][]int{{7}},
		},
		{
			name:   "single element used multiple times",
			nums:   []int{2},
			target: 6,
			want:   [][]int{{2, 2, 2}},
		},
		{
			name:   "target equals smallest element",
			nums:   []int{2, 3, 5},
			target: 2,
			want:   [][]int{{2}},
		},
		{
			name:   "all elements larger than target",
			nums:   []int{5, 10, 15},
			target: 3,
			want:   [][]int{},
		},
		{
			name:   "multiple paths with two elements",
			nums:   []int{2, 3},
			target: 6,
			want:   [][]int{{2, 2, 2}, {3, 3}},
		},
		{
			name:   "unsorted input",
			nums:   []int{9, 2, 6, 5},
			target: 9,
			want:   [][]int{{2, 2, 5}, {9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.nums, tt.target)
			if !equalSubsets(got, tt.want) {
				t.Fatalf("combinationSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
