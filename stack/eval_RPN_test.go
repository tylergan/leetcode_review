package stack

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "example 1",
			tokens: []string{"1", "2", "+", "3", "*", "4", "-"},
			want:   5,
		},
		{
			name:   "simple multiply",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "division with addition",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "long expression",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "division truncates toward zero",
			tokens: []string{"-3", "2", "/"},
			want:   -1,
		},
		{
			name:   "negative division",
			tokens: []string{"-4", "2", "/"},
			want:   -2,
		},
		{
			name:   "subtraction order matters",
			tokens: []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"},
			want:   14,
		},
		{
			name:   "single number",
			tokens: []string{"42"},
			want:   42,
		},
		{
			name:   "multiple negatives",
			tokens: []string{"-2", "3", "*", "4", "+", "5", "/"},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evalRPN(tt.tokens)
			if got != tt.want {
				t.Errorf("evalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
