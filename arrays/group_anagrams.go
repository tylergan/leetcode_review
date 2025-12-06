package arrays

import (
	"fmt"
)

/*
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:

Input: strs = ["x"]

Output: [["x"]]
Example 3:

Input: strs = [""]

Output: [[""]]
Constraints:

1 <= strs.length <= 1000.
0 <= strs[i].length <= 100
strs[i] is made up of lowercase English letters.
*/

func groupAnagrams(strs []string) [][]string {
	anagrams := map[string][]string{}
	for _, s := range strs {
		sHash := make([]int, 26)
		for _, c := range s {
			sHash[int(c)-int('a')]++
		}

		sHashStr := fmt.Sprintf("%v", sHash)
		anagrams[sHashStr] = append(anagrams[sHashStr], s)
	}

	var res [][]string
	for _, v := range anagrams {
		res = append(res, v)
	}

	return res
}
