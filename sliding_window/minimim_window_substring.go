package sliding_window

/*
Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".

You may assume that the correct output is always unique.

Example 1:

Input: s = "OUZODYXAZV", t = "XYZ"

Output: "YXAZ"
Explanation: "YXAZ" is the shortest substring that includes "X", "Y", and "Z" from string t.

Example 2:

Input: s = "xyz", t = "xyz"

Output: "xyz"
Example 3:

Input: s = "x", t = "xy"

Output: ""
Constraints:

1 <= s.length <= 1000
1 <= t.length <= 1000
s and t consist of uppercase and lowercase English letters.
*/

func minWindow(s string, t string) string {
	if s == t {
		return s
	}

	res, resLen := "", len(s)+1 // since it can include the entire string itself

	tCount := map[rune]int{}
	for _, c := range t {
		tCount[c]++
	}

	lettersComplete := 0
	currCount := map[rune]int{}
	for i, j := 0, 0; j < len(s); j++ {
		endChar := rune(s[j])
		if tFreq, ok := tCount[endChar]; ok {
			newCount := currCount[endChar] + 1 // increase our resources of t-only letters
			currCount[endChar] = newCount
			if newCount == tFreq { // if the counts are satisfied, increment letters complete
				lettersComplete++
			}
		}

		for lettersComplete == len(tCount) {
			if resLen > j-i+1 {
				resLen = j - i + 1
				res = s[i : j+1]
			}
			startChar := rune(s[i])
			if tFreq, ok := tCount[startChar]; ok {
				newCount := currCount[startChar] - 1 // decrement our t-only letter resource pool
				currCount[startChar] = newCount
				if newCount < tFreq { // moment our resources fall below the required count, no longer complete
					lettersComplete--
				}
			}
			i++
		}
	}

	return res
}
