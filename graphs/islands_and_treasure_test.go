package graphs

import (
	"slices"
	"testing"
)

const inf = 2147483647

func cloneIntMatrix(grid [][]int) [][]int {
	res := make([][]int, len(grid))
	for i := range grid {
		res[i] = slices.Clone(grid[i])
	}
	return res
}

func intMatricesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestIslandsAndTreasure(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			name: "example 1 - multiple treasure chests",
			grid: [][]int{
				{inf, -1, 0, inf},
				{inf, inf, inf, -1},
				{inf, -1, inf, -1},
				{0, -1, inf, inf},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "example 2 - one treasure",
			grid: [][]int{
				{0, -1},
				{inf, inf},
			},
			want: [][]int{
				{0, -1},
				{1, 2},
			},
		},
		{
			name: "unreachable land remains INF",
			grid: [][]int{
				{0, -1, inf},
				{-1, -1, inf},
				{inf, inf, inf},
			},
			want: [][]int{
				{0, -1, inf},
				{-1, -1, inf},
				{inf, inf, inf},
			},
		},
		{
			name: "all treasure chests",
			grid: [][]int{
				{0, 0},
				{0, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "empty land chooses nearest treasure",
			grid: [][]int{
				{0, inf, inf, 0},
			},
			want: [][]int{
				{0, 1, 1, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := cloneIntMatrix(tt.grid)
			islandsAndTreasure(grid)
			if !intMatricesEqual(grid, tt.want) {
				t.Fatalf("islandsAndTreasure() grid = %v, want %v", grid, tt.want)
			}
		})
	}
}
