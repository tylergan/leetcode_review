package arrays

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example 1: multiple anagram groups",
			strs: []string{"act", "pots", "tops", "cat", "stop", "hat"},
			want: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name: "example 2: single string",
			strs: []string{"x"},
			want: [][]string{{"x"}},
		},
		{
			name: "example 3: empty string",
			strs: []string{""},
			want: [][]string{{""}},
		},
		{
			name: "all same anagrams",
			strs: []string{"abc", "bca", "cab", "acb"},
			want: [][]string{{"abc", "bca", "cab", "acb"}},
		},
		{
			name: "no anagrams - all unique",
			strs: []string{"a", "b", "c", "d"},
			want: [][]string{{"a"}, {"b"}, {"c"}, {"d"}},
		},
		{
			name: "multiple empty strings",
			strs: []string{"", "", ""},
			want: [][]string{{"", "", ""}},
		},
		{
			name: "mixed lengths",
			strs: []string{"a", "aa", "aaa", "a"},
			want: [][]string{{"a", "a"}, {"aa"}, {"aaa"}},
		},
		{
			name: "different anagram groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "single character repeated",
			strs: []string{"a", "a", "a"},
			want: [][]string{{"a", "a", "a"}},
		},
		{
			name: "longer anagrams",
			strs: []string{"listen", "silent", "enlist", "hello", "world"},
			want: [][]string{{"listen", "silent", "enlist"}, {"hello"}, {"world"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !equalAnagramGroups(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalAnagramGroups compares two slices of string slices for equality,
// ignoring the order of groups and the order of strings within each group
func equalAnagramGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each group and convert to comparable format
	aSorted := make([]string, len(a))
	bSorted := make([]string, len(b))

	for i, group := range a {
		sorted := make([]string, len(group))
		copy(sorted, group)
		sort.Strings(sorted)
		aSorted[i] = stringSliceToString(sorted)
	}

	for i, group := range b {
		sorted := make([]string, len(group))
		copy(sorted, group)
		sort.Strings(sorted)
		bSorted[i] = stringSliceToString(sorted)
	}

	// Sort the groups themselves
	sort.Strings(aSorted)
	sort.Strings(bSorted)

	// Compare
	for i := range aSorted {
		if aSorted[i] != bSorted[i] {
			return false
		}
	}

	return true
}

// stringSliceToString converts a sorted slice of strings to a single string for comparison
func stringSliceToString(strs []string) string {
	result := ""
	for _, s := range strs {
		result += s + "|"
	}
	return result
}
