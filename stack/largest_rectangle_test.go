package stack

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "example 1",
			heights: []int{7, 1, 7, 2, 2, 4},
			want:    8,
		},
		{
			name:    "example 2",
			heights: []int{1, 3, 7},
			want:    7,
		},
		{
			name:    "classic",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "single bar",
			heights: []int{5},
			want:    5,
		},
		{
			name:    "all zeros",
			heights: []int{0, 0, 0},
			want:    0,
		},
		{
			name:    "increasing",
			heights: []int{1, 2, 3, 4, 5},
			want:    9,
		},
		{
			name:    "decreasing",
			heights: []int{5, 4, 3, 2, 1},
			want:    9,
		},
		{
			name:    "plateau",
			heights: []int{2, 2, 2, 2},
			want:    8,
		},
		{
			name:    "zero breaks width",
			heights: []int{2, 0, 2},
			want:    2,
		},
		{
			name:    "wide valley",
			heights: []int{6, 2, 5, 4, 5, 1, 6},
			want:    12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestRectangleArea(tt.heights)
			if got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}
