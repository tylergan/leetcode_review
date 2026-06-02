package graphs

import (
	"slices"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1 - all fruit eventually rots",
			grid: [][]int{
				{1, 1, 0},
				{0, 1, 1},
				{0, 1, 2},
			},
			want: 4,
		},
		{
			name: "example 2 - isolated fresh fruit",
			grid: [][]int{
				{1, 0, 1},
				{0, 2, 0},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			name: "no fresh fruit",
			grid: [][]int{
				{0, 2},
				{0, 0},
			},
			want: 0,
		},
		{
			name: "fresh fruit with no rotten source",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			want: -1,
		},
		{
			name: "multiple rotten sources meet in middle",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 1},
				{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "empty cells block spread",
			grid: [][]int{
				{2, 0, 1},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := orangesRotting(cloneIntMatrix(tt.grid))
			if got != tt.want {
				t.Fatalf("orangesRotting(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestOrangesRottingMutatesReachedFreshFruit(t *testing.T) {
	grid := [][]int{
		{2, 1},
		{1, 1},
	}

	got := orangesRotting(grid)
	if got != 2 {
		t.Fatalf("orangesRotting() = %d, want 2", got)
	}

	wantGrid := [][]int{
		{2, 2},
		{2, 2},
	}
	for i := range grid {
		if !slices.Equal(grid[i], wantGrid[i]) {
			t.Fatalf("grid row %d = %v, want %v", i, grid[i], wantGrid[i])
		}
	}
}
