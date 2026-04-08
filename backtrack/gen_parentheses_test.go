package backtrack

import (
	"fmt"
	"slices"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n=1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n=2",
			n:    2,
			want: []string{"(())", "()()" },
		},
		{
			name: "n=3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()" },
		},
		{
			name: "n=4",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
				"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
				"()(()())", "()(())()", "()()(())", "()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			slices.Sort(got)
			slices.Sort(tt.want)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("generateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesisCount(t *testing.T) {
	// Catalan numbers: C(n) = (2n)! / ((n+1)! * n!)
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 2},
		{n: 3, want: 5},
		{n: 4, want: 14},
		{n: 5, want: 42},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			got := len(generateParenthesis(tt.n))
			if got != tt.want {
				t.Fatalf("len(generateParenthesis(%d)) = %d, want %d (Catalan number)", tt.n, got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesisAllValid(t *testing.T) {
	for n := 1; n <= 5; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			results := generateParenthesis(n)
			for _, s := range results {
				if len(s) != 2*n {
					t.Fatalf("generateParenthesis(%d) produced %q with length %d, want %d", n, s, len(s), 2*n)
				}
				depth := 0
				for _, ch := range s {
					if ch == '(' {
						depth++
					} else if ch == ')' {
						depth--
					} else {
						t.Fatalf("unexpected character %q in %q", string(ch), s)
					}
					if depth < 0 {
						t.Fatalf("invalid parentheses %q: closing paren without matching open", s)
					}
				}
				if depth != 0 {
					t.Fatalf("invalid parentheses %q: %d unclosed open parens", s, depth)
				}
			}
		})
	}
}

func TestGenerateParenthesisNoDuplicates(t *testing.T) {
	for n := 1; n <= 5; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			results := generateParenthesis(n)
			seen := map[string]bool{}
			for _, s := range results {
				if seen[s] {
					t.Fatalf("duplicate result %q in generateParenthesis(%d)", s, n)
				}
				seen[s] = true
			}
		})
	}
}
