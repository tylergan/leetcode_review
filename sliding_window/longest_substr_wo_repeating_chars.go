package sliding_window

/*
Given a string s, find the length of the longest substring without duplicate characters.

A substring is a contiguous sequence of characters within a string.

Example 1:

Input: s = "zxyzxyz"

Output: 3
Explanation: The string "xyz" is the longest without duplicate characters.

Example 2:

Input: s = "xxxx"

Output: 1
Constraints:

0 <= s.length <= 1000
s may consist of printable ASCII characters.

*/

func lengthOfLongestSubstring(s string) int {
	res := 0
	inSubstr := map[rune]bool{}

	i := 0
	for j, char := range s {
		for i < j && inSubstr[char] {
			delete(inSubstr, rune(s[i]))
			i++
		}
		inSubstr[char] = true
		res = max(res, j-i+1)
	}

	return res
}
