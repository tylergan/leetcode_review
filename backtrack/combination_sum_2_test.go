package backtrack

import (
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "example 1 - duplicates in input",
			candidates: []int{9, 2, 2, 4, 6, 1, 5},
			target:     8,
			want:       [][]int{{1, 2, 5}, {2, 2, 4}, {2, 6}},
		},
		{
			name:       "example 2 - no duplicates in input",
			candidates: []int{1, 2, 3, 4, 5},
			target:     7,
			want:       [][]int{{1, 2, 4}, {2, 5}, {3, 4}},
		},
		{
			name:       "all duplicates - only one valid combination",
			candidates: []int{1, 1, 1},
			target:     2,
			want:       [][]int{{1, 1}},
		},
		{
			name:       "duplicates do not produce duplicate combinations",
			candidates: []int{2, 2, 2},
			target:     4,
			want:       [][]int{{2, 2}},
		},
		{
			name:       "single element equals target",
			candidates: []int{7},
			target:     7,
			want:       [][]int{{7}},
		},
		{
			name:       "element cannot be reused",
			candidates: []int{4},
			target:     8,
			want:       [][]int{},
		},
		{
			name:       "no valid combination",
			candidates: []int{5, 10, 15},
			target:     3,
			want:       [][]int{},
		},
		{
			name:       "target is minimum value 1",
			candidates: []int{1, 2, 3},
			target:     1,
			want:       [][]int{{1}},
		},
		{
			name:       "each element used at most once",
			candidates: []int{1, 2, 3},
			target:     6,
			want:       [][]int{{1, 2, 3}},
		},
		{
			name:       "multiple duplicates with multiple valid combinations",
			candidates: []int{1, 1, 2, 5, 6, 7, 10},
			target:     8,
			want:       [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.candidates, tt.target)
			if !equalSubsets(got, tt.want) {
				t.Fatalf("combinationSum2(%v, %d) = %v, want %v", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}
