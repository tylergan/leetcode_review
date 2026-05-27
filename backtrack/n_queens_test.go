package backtrack

import (
	"fmt"
	"slices"
	"testing"
)

func equalBoards(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	sortBoards(a)
	sortBoards(b)

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func sortBoards(boards [][]string) {
	slices.SortFunc(boards, func(a, b []string) int {
		for i := range min(len(a), len(b)) {
			if a[i] != b[i] {
				if a[i] < b[i] {
					return -1
				}
				return 1
			}
		}
		return len(a) - len(b)
	})
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{
			name: "example 1 - n=4 has two solutions",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			name: "example 2 - n=1",
			n:    1,
			want: [][]string{{"Q"}},
		},
		{
			name: "n=2 has no solutions",
			n:    2,
			want: [][]string{},
		},
		{
			name: "n=3 has no solutions",
			n:    3,
			want: [][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.n)
			if !equalBoards(got, tt.want) {
				t.Fatalf("solveNQueens(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestSolveNQueensCounts(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 0},
		{n: 3, want: 0},
		{n: 4, want: 2},
		{n: 5, want: 10},
		{n: 6, want: 4},
		{n: 7, want: 40},
		{n: 8, want: 92},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			got := solveNQueens(tt.n)
			if len(got) != tt.want {
				t.Fatalf("len(solveNQueens(%d)) = %d, want %d", tt.n, len(got), tt.want)
			}
		})
	}
}

func TestSolveNQueensBoardsAreValid(t *testing.T) {
	for n := 1; n <= 8; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			seen := map[string]bool{}
			for _, board := range solveNQueens(n) {
				key := fmt.Sprint(board)
				if seen[key] {
					t.Fatalf("solveNQueens(%d) contains duplicate board %v", n, board)
				}
				seen[key] = true

				assertValidNQueensBoard(t, board)
			}
		})
	}
}

func assertValidNQueensBoard(t *testing.T, board []string) {
	t.Helper()

	n := len(board)
	rows := map[int]bool{}
	cols := map[int]bool{}
	forwardDiags := map[int]bool{}
	backwardDiags := map[int]bool{}
	queens := 0

	for r, row := range board {
		if len(row) != n {
			t.Fatalf("row %q has length %d, want %d", row, len(row), n)
		}

		for c, cell := range row {
			switch cell {
			case 'Q':
				queens++
				if rows[r] || cols[c] || forwardDiags[r+c] || backwardDiags[r-c] {
					t.Fatalf("board has attacking queens: %v", board)
				}
				rows[r] = true
				cols[c] = true
				forwardDiags[r+c] = true
				backwardDiags[r-c] = true
			case '.':
			default:
				t.Fatalf("board contains unexpected cell %q: %v", cell, board)
			}
		}
	}

	if queens != n {
		t.Fatalf("board has %d queens, want %d: %v", queens, n, board)
	}
}
