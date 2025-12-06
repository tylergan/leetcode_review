package arrays

/*
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false
Constraints:

s and t consist of lowercase English letters.

*/

import (
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sHash, tHash := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s); i++ {
		sHash[int(s[i])-int('a')] += 1
		tHash[int(t[i])-int('a')] += 1
	}

	return slices.Equal(sHash, tHash)
}
