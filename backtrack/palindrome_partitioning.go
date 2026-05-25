package backtrack

import (
	"slices"
	"strings"
)

/*
Given a string s, split s into substrings where every substring is a palindrome. Return all possible lists of palindromic substrings.

You may return the solution in any order.

Example 1:

Input: s = "aab"

Output: [["a","a","b"],["aa","b"]]
Example 2:

Input: s = "a"

Output: [["a"]]
Constraints:

1 <= s.length <= 20
s contains only lowercase English letters.
*/

func partition(s string) [][]string {
	var res [][]string
	curr := []string{}

	var dfs func(string)
	dfs = func(s string) {
		if s == "" {
			res = append(res, slices.Clone(curr))
		}

		for i := 0; i < len(s); i++ {
			substr := s[:i+1]

			if !isPal(substr) {
				continue
			}

			curr = append(curr, substr)
			dfs(s[i+1:])
			curr = curr[:len(curr)-1]
		}
	}

	dfs(s)

	return res
}

func isPal(s string) bool {
	n := len(s)
	var sb strings.Builder

	for i := n - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}

	return s == sb.String()
}
