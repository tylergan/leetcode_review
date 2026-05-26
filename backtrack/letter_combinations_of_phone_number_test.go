package backtrack

import (
	"fmt"
	"slices"
	"testing"
)

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	slices.Sort(a)
	slices.Sort(b)

	return slices.Equal(a, b)
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "example 1 - two digits",
			digits: "34",
			want:   []string{"dg", "dh", "di", "eg", "eh", "ei", "fg", "fh", "fi"},
		},
		{
			name:   "example 2 - empty input",
			digits: "",
			want:   []string{},
		},
		{
			name:   "single digit with three letters",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "single digit with four letters",
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			name:   "two digits include four-letter branch",
			digits: "79",
			want: []string{
				"pw", "px", "py", "pz",
				"qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz",
				"sw", "sx", "sy", "sz",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			if !equalStringSlices(got, tt.want) {
				t.Fatalf("letterCombinations(%q) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}

func TestLetterCombinationsCount(t *testing.T) {
	tests := []struct {
		digits string
		want   int
	}{
		{digits: "", want: 0},
		{digits: "2", want: 3},
		{digits: "22", want: 9},
		{digits: "27", want: 12},
		{digits: "2799", want: 192},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("digits=%q", tt.digits), func(t *testing.T) {
			got := letterCombinations(tt.digits)
			if len(got) != tt.want {
				t.Fatalf("len(letterCombinations(%q)) = %d, want %d", tt.digits, len(got), tt.want)
			}
		})
	}
}

func TestLetterCombinationsNoDuplicates(t *testing.T) {
	for _, digits := range []string{"2", "34", "79", "2799"} {
		t.Run(digits, func(t *testing.T) {
			got := letterCombinations(digits)
			seen := map[string]bool{}
			for _, combination := range got {
				if seen[combination] {
					t.Fatalf("letterCombinations(%q) contains duplicate %q in %v", digits, combination, got)
				}
				seen[combination] = true
			}
		})
	}
}
