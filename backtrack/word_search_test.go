package backtrack

import (
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name: "example 1 - word present",
			board: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'S', 'A', 'A', 'T'},
				{'A', 'C', 'A', 'E'},
			},
			word: "CAT",
			want: true,
		},
		{
			name: "example 2 - word not present",
			board: [][]byte{
				{'A', 'B', 'C', 'D'},
				{'S', 'A', 'A', 'T'},
				{'A', 'C', 'A', 'E'},
			},
			word: "BAT",
			want: false,
		},
		{
			name: "single cell match",
			board: [][]byte{
				{'A'},
			},
			word: "A",
			want: true,
		},
		{
			name: "single cell no match",
			board: [][]byte{
				{'A'},
			},
			word: "B",
			want: false,
		},
		{
			name: "word longer than board cells",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "ABCDA",
			want: false,
		},
		{
			name: "same cell cannot be reused",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "ABA",
			want: false,
		},
		{
			name: "snake path through grid",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
			want: true,
		},
		{
			name: "path requires backtracking",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "SEE",
			want: true,
		},
		{
			name: "word not possible due to reuse",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
			want: false,
		},
		{
			name: "full board traversal",
			board: [][]byte{
				{'A', 'B'},
				{'D', 'C'},
			},
			word: "ABCD",
			want: true,
		},
		{
			name: "single row",
			board: [][]byte{
				{'A', 'B', 'C', 'D'},
			},
			word: "BCD",
			want: true,
		},
		{
			name: "single column",
			board: [][]byte{
				{'A'},
				{'B'},
				{'C'},
			},
			word: "ABC",
			want: true,
		},
		{
			name: "all same characters - valid path exists",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
				{'A', 'A', 'A'},
			},
			word: "AAAAAAAAA",
			want: true,
		},
		{
			name: "all same characters - word too long",
			board: [][]byte{
				{'A', 'A'},
				{'A', 'A'},
			},
			word: "AAAAA",
			want: false,
		},
		{
			name: "start from last cell",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'A'},
				{'X', 'X', 'B'},
			},
			word: "AB",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Fatalf("exist(%v, %q) = %v, want %v", tt.board, tt.word, got, tt.want)
			}
		})
	}
}
