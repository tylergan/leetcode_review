package sliding_window

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "OUZODYXAZV",
			t:    "XYZ",
			want: "YXAZ",
		},
		{
			name: "example 2",
			s:    "xyz",
			t:    "xyz",
			want: "xyz",
		},
		{
			name: "example 3",
			s:    "x",
			t:    "xy",
			want: "",
		},
		{
			name: "classic",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "single character",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "t longer than s",
			s:    "ab",
			t:    "abc",
			want: "",
		},
		{
			name: "duplicates required",
			s:    "aaabdabcefaecbef",
			t:    "aabc",
			want: "abdabc",
		},
		{
			name: "case sensitivity",
			s:    "aAaA",
			t:    "Aa",
			want: "aA",
		},
		{
			name: "window at end",
			s:    "zzzzzzabc",
			t:    "abc",
			want: "abc",
		},
		{
			name: "no possible window",
			s:    "abcdef",
			t:    "xyz",
			want: "",
		},
		{
			name: "empty s",
			s:    "",
			t:    "a",
			want: "",
		},
		{
			name: "all same required",
			s:    "bba",
			t:    "bb",
			want: "bb",
		},
		{
			name: "single letter",
			s:    "ab",
			t:    "a",
			want: "a",
		},
		{
			name: "another simple test",
			s:    "abc",
			t:    "ac",
			want: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
