package sliding_window

/*
You are given a string s consisting of only uppercase english characters and an integer k. You can choose up to k characters of the string and replace them with any other uppercase English character.

After performing at most k replacements, return the length of the longest substring which contains only one distinct character.

Example 1:

Input: s = "XYYX", k = 2

Output: 4
Explanation: Either replace the 'X's with 'Y's, or replace the 'Y's with 'X's.

Example 2:

Input: s = "AAABABB", k = 1

Output: 5
Constraints:

1 <= s.length <= 1000
0 <= k <= s.length
*/

func characterReplacement(s string, k int) int {
	res := 0
	count := [26]int{}
	maxFreq := 0
	for i, j := 0, 0; j < len(s); j++ {
		// Update the max character frequency
		currChar := s[j] - 'A'
		count[currChar]++
		maxFreq = max(maxFreq, count[currChar])

		// Characters that do not have the highest freq
		// are the ones we want to replace (which is just
		// the len(substring) - maxFreq). If the # replaced
		// greater than k, we need to shorten the substring
		for i < j && (j-i+1)-maxFreq > k {
			count[s[i]-'A']--
			i++
		}

		res = max(res, j-i+1)
	}

	return res
}
