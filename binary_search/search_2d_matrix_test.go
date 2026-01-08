package binary_search

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "example 1: target present",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target: 10,
			want:   true,
		},
		{
			name: "example 2: target missing",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target: 15,
			want:   false,
		},
		{
			name: "single row target present",
			matrix: [][]int{
				{-5, -2, 0, 3, 9},
			},
			target: 3,
			want:   true,
		},
		{
			name: "single row target missing",
			matrix: [][]int{
				{-5, -2, 0, 3, 9},
			},
			target: 4,
			want:   false,
		},
		{
			name: "single column target present",
			matrix: [][]int{
				{1},
				{4},
				{7},
			},
			target: 7,
			want:   true,
		},
		{
			name: "single column target missing",
			matrix: [][]int{
				{1},
				{4},
				{7},
			},
			target: 6,
			want:   false,
		},
		{
			name: "target is first element",
			matrix: [][]int{
				{2, 3, 5},
				{7, 11, 13},
			},
			target: 2,
			want:   true,
		},
		{
			name: "target is last element",
			matrix: [][]int{
				{2, 3, 5},
				{7, 11, 13},
			},
			target: 13,
			want:   true,
		},
		{
			name: "target between rows",
			matrix: [][]int{
				{1, 2, 3},
				{10, 11, 12},
				{20, 21, 22},
			},
			target: 9,
			want:   false,
		},
		{
			name: "target less than minimum",
			matrix: [][]int{
				{1, 2, 3},
				{10, 11, 12},
			},
			target: 0,
			want:   false,
		},
		{
			name: "target greater than maximum",
			matrix: [][]int{
				{1, 2, 3},
				{10, 11, 12},
			},
			target: 13,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Fatalf("searchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}
