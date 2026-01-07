package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "example 1",
			s:    "[]",
			want: true,
		},
		{
			name: "example 2",
			s:    "([{}])",
			want: true,
		},
		{
			name: "example 3",
			s:    "[(])",
			want: false,
		},
		{
			name: "single open",
			s:    "(",
			want: false,
		},
		{
			name: "single close",
			s:    ")",
			want: false,
		},
		{
			name: "simple pair",
			s:    "()",
			want: true,
		},
		{
			name: "nested mixed",
			s:    "{[()]}",
			want: true,
		},
		{
			name: "wrong order",
			s:    "([)]",
			want: false,
		},
		{
			name: "extra open",
			s:    "((()))(",
			want: false,
		},
		{
			name: "extra close",
			s:    "((()))]",
			want: false,
		},
		{
			name: "multiple pairs",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "interleaved invalid",
			s:    "([]{})]",
			want: false,
		},
		{
			name: "only opens",
			s:    "([{{[",
			want: false,
		},
		{
			name: "only closes",
			s:    ")]}]])",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
