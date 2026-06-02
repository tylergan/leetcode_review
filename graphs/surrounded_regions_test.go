package graphs

import (
	"slices"
	"testing"
)

func cloneByteBoard(board [][]byte) [][]byte {
	res := make([][]byte, len(board))
	for i := range board {
		res[i] = slices.Clone(board[i])
	}
	return res
}

func boardsEqualBytes(a, b [][]byte) bool {
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

func TestSolveSurroundedRegions(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{
			name: "example 1 - surrounded center flips but border stays",
			board: [][]byte{
				[]byte("XXXX"),
				[]byte("XOOX"),
				[]byte("XOOX"),
				[]byte("XXXO"),
			},
			want: [][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXO"),
			},
		},
		{
			name: "border-connected snake remains safe",
			board: [][]byte{
				[]byte("OXXOX"),
				[]byte("XOOXO"),
				[]byte("XOXOX"),
				[]byte("OXOOO"),
				[]byte("XXOXO"),
			},
			want: [][]byte{
				[]byte("OXXOX"),
				[]byte("XXXXO"),
				[]byte("XXXOX"),
				[]byte("OXOOO"),
				[]byte("XXOXO"),
			},
		},
		{
			name: "diagonal border connection does not save region",
			board: [][]byte{
				[]byte("OXX"),
				[]byte("XOX"),
				[]byte("XXX"),
			},
			want: [][]byte{
				[]byte("OXX"),
				[]byte("XXX"),
				[]byte("XXX"),
			},
		},
		{
			name: "all border-connected region remains",
			board: [][]byte{
				[]byte("OOO"),
				[]byte("OOO"),
				[]byte("OOO"),
			},
			want: [][]byte{
				[]byte("OOO"),
				[]byte("OOO"),
				[]byte("OOO"),
			},
		},
		{
			name: "single cell O is border and remains",
			board: [][]byte{
				[]byte("O"),
			},
			want: [][]byte{
				[]byte("O"),
			},
		},
		{
			name: "single cell X remains",
			board: [][]byte{
				[]byte("X"),
			},
			want: [][]byte{
				[]byte("X"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := cloneByteBoard(tt.board)
			solve(board)
			if !boardsEqualBytes(board, tt.want) {
				t.Fatalf("solve() board = %q, want %q", board, tt.want)
			}
		})
	}
}
