package two_pointers

import (
	"strings"
	"unicode"
)

/*
Given a string s, return true if it is a palindrome, otherwise return false.

A palindrome is a string that reads the same forward and backward. It is also case-insensitive and ignores all non-alphanumeric characters.

Note: Alphanumeric characters consist of letters (A-Z, a-z) and numbers (0-9).

Example 1:

Input: s = "Was it a car or a cat I saw?"

Output: true
Explanation: After considering only alphanumerical characters we have "wasitacaroracatisaw", which is a palindrome.

Example 2:

Input: s = "tab a cat"

Output: false
Explanation: "tabacat" is not a palindrome.

Constraints:

1 <= s.length <= 1000
s is made up of only printable ASCII characters.
*/

func isPalindrome(s string) bool {
	isAlphaNum := func(char rune) bool {
		return unicode.IsLetter(char) || unicode.IsNumber(char)
	}

	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlphaNum(rune(s[i])) {
			i++
		}
		for i < j && !isAlphaNum(rune(s[j])) {
			j--
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}
