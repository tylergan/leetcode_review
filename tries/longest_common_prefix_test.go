package tries

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "example 1 - shared two char prefix",
			strs: []string{"bat", "bag", "bank", "band"},
			want: "ba",
		},
		{
			name: "example 2 - branching after prefix",
			strs: []string{"dance", "dag", "danger", "damage"},
			want: "da",
		},
		{
			name: "example 3 - no common prefix",
			strs: []string{"neet", "feet"},
			want: "",
		},
		{
			name: "single string returns itself",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "identical strings",
			strs: []string{"same", "same", "same"},
			want: "same",
		},
		{
			name: "one string is the prefix of the others",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "prefix is an entire shorter word",
			strs: []string{"interspecies", "interstellar", "inter"},
			want: "inter",
		},
		{
			name: "empty string in the middle forces empty prefix",
			strs: []string{"abc", "", "abd"},
			want: "",
		},
		{
			name: "first string empty",
			strs: []string{"", "abc"},
			want: "",
		},
		{
			name: "order does not change the result",
			strs: []string{"flight", "flower", "flow"},
			want: "fl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Fatalf("longestCommonPrefix(%q) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
