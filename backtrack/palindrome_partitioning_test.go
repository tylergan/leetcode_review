package backtrack

import (
	"slices"
	"testing"
)

func equalStringPartitions(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	sortStringPartitions(a)
	sortStringPartitions(b)

	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func sortStringPartitions(partitions [][]string) {
	slices.SortFunc(partitions, func(a, b []string) int {
		for i := range min(len(a), len(b)) {
			if a[i] != b[i] {
				if a[i] < b[i] {
					return -1
				}
				return 1
			}
		}
		return len(a) - len(b)
	})
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want [][]string
	}{
		{
			name: "example 1 - split or group leading palindrome",
			s:    "aab",
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "example 2 - single character",
			s:    "a",
			want: [][]string{{"a"}},
		},
		{
			name: "all single characters when no longer palindrome exists",
			s:    "abc",
			want: [][]string{{"a", "b", "c"}},
		},
		{
			name: "whole string palindrome plus smaller partitions",
			s:    "aba",
			want: [][]string{{"a", "b", "a"}, {"aba"}},
		},
		{
			name: "repeated characters generate every cut combination",
			s:    "aaa",
			want: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "palindrome can appear in the middle",
			s:    "abbc",
			want: [][]string{{"a", "b", "b", "c"}, {"a", "bb", "c"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.s)
			if !equalStringPartitions(got, tt.want) {
				t.Fatalf("partition(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestIsPal(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "a", want: true},
		{s: "aa", want: true},
		{s: "aba", want: true},
		{s: "abba", want: true},
		{s: "ab", want: false},
		{s: "abc", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := isPal(tt.s)
			if got != tt.want {
				t.Fatalf("isPal(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
