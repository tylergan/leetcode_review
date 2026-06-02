package graphs

import (
	"slices"
	"testing"
)

func equalCoordinates(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sortCoordinates(a)
	sortCoordinates(b)

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func sortCoordinates(coords [][]int) {
	slices.SortFunc(coords, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "example 1 - meeting from both oceans",
			heights: [][]int{
				{4, 2, 7, 3, 4},
				{7, 4, 6, 4, 7},
				{6, 3, 5, 3, 6},
			},
			want: [][]int{
				{0, 2},
				{0, 4},
				{1, 0},
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 0},
			},
		},
		{
			name: "example 2 - single column",
			heights: [][]int{
				{1},
				{1},
			},
			want: [][]int{{0, 0}, {1, 0}},
		},
		{
			name: "single cell touches both oceans",
			heights: [][]int{
				{5},
			},
			want: [][]int{{0, 0}},
		},
		{
			name: "flat grid all cells reach both oceans",
			heights: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "single row all cells touch both oceans",
			heights: [][]int{
				{3, 1, 2},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pacificAtlantic(tt.heights)
			if !equalCoordinates(got, tt.want) {
				t.Fatalf("pacificAtlantic(%v) = %v, want %v", tt.heights, got, tt.want)
			}
		})
	}
}
