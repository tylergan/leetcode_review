package sliding_window

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "example 1",
			s1:   "abc",
			s2:   "lecabee",
			want: true,
		},
		{
			name: "example 2",
			s1:   "abc",
			s2:   "lecaabee",
			want: false,
		},
		{
			name: "s1 longer than s2",
			s1:   "abcd",
			s2:   "abc",
			want: false,
		},
		{
			name: "exact match",
			s1:   "abc",
			s2:   "abc",
			want: true,
		},
		{
			name: "permutation at start",
			s1:   "ab",
			s2:   "baaaa",
			want: true,
		},
		{
			name: "permutation in middle",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "permutation at end",
			s1:   "ab",
			s2:   "eidoooba",
			want: true,
		},
		{
			name: "repeated chars need counts",
			s1:   "aabc",
			s2:   "caaebaa",
			want: false,
		},
		{
			name: "repeated chars match",
			s1:   "aabc",
			s2:   "zzcaabyy",
			want: true,
		},
		{
			name: "single char",
			s1:   "a",
			s2:   "bbbba",
			want: true,
		},
		{
			name: "single char missing",
			s1:   "a",
			s2:   "bbbb",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.want {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}
