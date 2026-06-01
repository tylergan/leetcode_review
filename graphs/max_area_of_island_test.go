package graphs

import (
	"slices"
	"testing"
)

func cloneIntGrid(grid [][]int) [][]int {
	res := make([][]int, len(grid))
	for i := range grid {
		res[i] = slices.Clone(grid[i])
	}
	return res
}

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1 - diagonal land does not connect",
			grid: [][]int{
				{0, 1, 1, 0, 1},
				{1, 0, 1, 0, 1},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 1},
			},
			want: 6,
		},
		{
			name: "all water",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
			},
			want: 6,
		},
		{
			name: "single cell land",
			grid: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]int{
				{0},
			},
			want: 0,
		},
		{
			name: "chooses maximum among many islands",
			grid: [][]int{
				{1, 1, 0, 0, 1},
				{1, 0, 0, 1, 1},
				{0, 0, 1, 0, 0},
				{1, 1, 1, 0, 1},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAreaOfIsland(cloneIntGrid(tt.grid))
			if got != tt.want {
				t.Fatalf("maxAreaOfIsland(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestMaxAreaOfIslandSinksVisitedLand(t *testing.T) {
	grid := [][]int{
		{1, 0, 1},
		{1, 1, 1},
	}

	got := maxAreaOfIsland(grid)
	if got != 5 {
		t.Fatalf("maxAreaOfIsland() = %d, want 5", got)
	}

	wantGrid := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}
	for i := range grid {
		if !slices.Equal(grid[i], wantGrid[i]) {
			t.Fatalf("grid row %d = %v, want %v", i, grid[i], wantGrid[i])
		}
	}
}
