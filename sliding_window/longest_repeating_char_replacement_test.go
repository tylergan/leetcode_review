package sliding_window

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "example 1",
			s:    "XYYX",
			k:    2,
			want: 4,
		},
		{
			name: "example 2",
			s:    "AAABABB",
			k:    1,
			want: 5,
		},
		{
			name: "single char",
			s:    "A",
			k:    0,
			want: 1,
		},
		{
			name: "all same",
			s:    "AAAA",
			k:    2,
			want: 4,
		},
		{
			name: "no replacements",
			s:    "ABAB",
			k:    0,
			want: 1,
		},
		{
			name: "k allows full string",
			s:    "ABCD",
			k:    3,
			want: 4,
		},
		{
			name: "classic case",
			s:    "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "window shifts",
			s:    "ABBB",
			k:    2,
			want: 4,
		},
		{
			name: "k zero with runs",
			s:    "AABBA",
			k:    0,
			want: 2,
		},
		{
			name: "large k",
			s:    "ABCDE",
			k:    10,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.want {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}
