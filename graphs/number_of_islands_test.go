package graphs

import (
	"slices"
	"testing"
)

func cloneGrid(grid [][]byte) [][]byte {
	res := make([][]byte, len(grid))
	for i := range grid {
		res[i] = slices.Clone(grid[i])
	}
	return res
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example 1 - connected island",
			grid: [][]byte{
				[]byte("01110"),
				[]byte("01010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			want: 1,
		},
		{
			name: "example 2 - four separate islands",
			grid: [][]byte{
				[]byte("11001"),
				[]byte("11001"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 4,
		},
		{
			name: "all water",
			grid: [][]byte{
				[]byte("000"),
				[]byte("000"),
			},
			want: 0,
		},
		{
			name: "all land is one island",
			grid: [][]byte{
				[]byte("111"),
				[]byte("111"),
			},
			want: 1,
		},
		{
			name: "diagonal land does not connect",
			grid: [][]byte{
				[]byte("101"),
				[]byte("010"),
				[]byte("101"),
			},
			want: 5,
		},
		{
			name: "single cell land",
			grid: [][]byte{
				[]byte("1"),
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{
				[]byte("0"),
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIslands(cloneGrid(tt.grid))
			if got != tt.want {
				t.Fatalf("numIslands(%q) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}

func TestNumIslandsSinksVisitedLand(t *testing.T) {
	grid := [][]byte{
		[]byte("101"),
		[]byte("111"),
	}

	got := numIslands(grid)
	if got != 1 {
		t.Fatalf("numIslands() = %d, want 1", got)
	}

	wantGrid := [][]byte{
		[]byte("000"),
		[]byte("000"),
	}
	for i := range grid {
		if !slices.Equal(grid[i], wantGrid[i]) {
			t.Fatalf("grid row %d = %q, want %q", i, grid[i], wantGrid[i])
		}
	}
}
