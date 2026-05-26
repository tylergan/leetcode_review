package backtrack

/*
You are given a string digits made up of digits from 2 through 9 inclusive.

Each digit (not including 1) is mapped to a set of characters as shown below:

A digit could represent any one of the characters it maps to.

Return all possible letter combinations that digits could represent. You may return the answer in any order.

Example 1:

Input: digits = "34"

Output: ["dg","dh","di","eg","eh","ei","fg","fh","fi"]
Example 2:

Input: digits = ""

Output: []
Constraints:

0 <= digits.length <= 4
2 <= digits[i] <= 9
*/

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapping := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	var res []string
	curr := ""

	var dfs func(int)
	dfs = func(dIdx int) {
		if dIdx == len(digits) {
			res = append(res, curr)
			return
		}

		currDigit := string(digits[dIdx])
		letters := mapping[currDigit]

		for i := 0; i < len(letters); i++ {
			curr = curr + letters[i]
			dfs(dIdx + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)

	return res
}
