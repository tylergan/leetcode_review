package tries

import (
	"slices"
	"testing"
)

func equalWordResults(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}

func cloneBoard(board [][]byte) [][]byte {
	res := make([][]byte, len(board))
	for i := range board {
		res[i] = slices.Clone(board[i])
	}
	return res
}

func boardsEqual(a, b [][]byte) bool {
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

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{
			name: "example 1 - multiple words share paths",
			board: [][]byte{
				[]byte("abcd"),
				[]byte("saat"),
				[]byte("acke"),
				[]byte("acdn"),
			},
			words: []string{"bat", "cat", "back", "backend", "stack"},
			want:  []string{"cat", "back", "backend"},
		},
		{
			name: "example 2 - cannot reuse cells",
			board: [][]byte{
				[]byte("xo"),
				[]byte("xo"),
			},
			words: []string{"xoxo"},
			want:  []string{},
		},
		{
			name: "single cell",
			board: [][]byte{
				[]byte("a"),
			},
			words: []string{"a", "b", "aa"},
			want:  []string{"a"},
		},
		{
			name: "prefix words can both be found",
			board: [][]byte{
				[]byte("ab"),
				[]byte("cd"),
			},
			words: []string{"a", "ab", "abc", "abd", "acdb"},
			want:  []string{"a", "ab", "abd", "acdb"},
		},
		{
			name: "same word discovered from multiple starts only appears once",
			board: [][]byte{
				[]byte("aa"),
				[]byte("aa"),
			},
			words: []string{"aa", "aaa", "aaaa"},
			want:  []string{"aa", "aaa", "aaaa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := cloneBoard(tt.board)
			got := findWords(board, tt.words)
			if !equalWordResults(got, tt.want) {
				t.Fatalf("findWords(%v, %v) = %v, want %v", tt.board, tt.words, got, tt.want)
			}
			if !boardsEqual(board, tt.board) {
				t.Fatalf("findWords mutated board: got %v, want %v", board, tt.board)
			}
		})
	}
}
