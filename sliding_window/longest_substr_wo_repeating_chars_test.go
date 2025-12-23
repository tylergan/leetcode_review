package sliding_window

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 1: repeating pattern",
			s:    "zxyzxyz",
			want: 3,
		},
		{
			name: "example 2: all same",
			s:    "xxxx",
			want: 1,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "all unique",
			s:    "abcdef",
			want: 6,
		},
		{
			name: "classic example",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "overlap reset",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "repeated middle",
			s:    "abba",
			want: 2,
		},
		{
			name: "space char",
			s:    " ",
			want: 1,
		},
		{
			name: "late duplicate",
			s:    "dvdf",
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
