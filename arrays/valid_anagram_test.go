package arrays

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		tStr     string
		expected bool
	}{
		{
			name:     "example 1 - racecar and carrace",
			s:        "racecar",
			tStr:     "carrace",
			expected: true,
		},
		{
			name:     "example 2 - jar and jam",
			s:        "jar",
			tStr:     "jam",
			expected: false,
		},
		{
			name:     "empty strings",
			s:        "",
			tStr:     "",
			expected: true,
		},
		{
			name:     "single character same",
			s:        "a",
			tStr:     "a",
			expected: true,
		},
		{
			name:     "single character different",
			s:        "a",
			tStr:     "b",
			expected: false,
		},
		{
			name:     "different lengths - first longer",
			s:        "abc",
			tStr:     "ab",
			expected: false,
		},
		{
			name:     "different lengths - second longer",
			s:        "ab",
			tStr:     "abc",
			expected: false,
		},
		{
			name:     "same string",
			s:        "listen",
			tStr:     "listen",
			expected: true,
		},
		{
			name:     "valid anagram - listen and silent",
			s:        "listen",
			tStr:     "silent",
			expected: true,
		},
		{
			name:     "valid anagram - anagram and nagaram",
			s:        "anagram",
			tStr:     "nagaram",
			expected: true,
		},
		{
			name:     "not an anagram - different characters",
			s:        "rat",
			tStr:     "car",
			expected: false,
		},
		{
			name:     "not an anagram - same length different freq",
			s:        "aab",
			tStr:     "abb",
			expected: false,
		},
		{
			name:     "repeated characters - valid anagram",
			s:        "aabbcc",
			tStr:     "abcabc",
			expected: true,
		},
		{
			name:     "repeated characters - invalid anagram",
			s:        "aaabbb",
			tStr:     "aabbbb",
			expected: false,
		},
		{
			name:     "all same character - valid",
			s:        "aaaa",
			tStr:     "aaaa",
			expected: true,
		},
		{
			name:     "all same character - different count",
			s:        "aaa",
			tStr:     "aaaa",
			expected: false,
		},
		{
			name:     "alphabet scrambled",
			s:        "abcdefghijklmnopqrstuvwxyz",
			tStr:     "zyxwvutsrqponmlkjihgfedcba",
			expected: true,
		},
		{
			name:     "long string - valid anagram",
			s:        "thequickbrownfoxjumpsoverthelazydog",
			tStr:     "ogdalyztherevospmujxofnworbukciqeht",
			expected: true,
		},
		{
			name:     "long string - invalid anagram",
			s:        "thequickbrownfoxjumpsoverthelazydog",
			tStr:     "thequickbrownfoxjumpsoveraplazydog",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isAnagram(tt.s, tt.tStr)
			if result != tt.expected {
				t.Errorf("isAnagram(%q, %q) = %v; expected %v", tt.s, tt.tStr, result, tt.expected)
			}
		})
	}
}
