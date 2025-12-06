package arrays

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Design an algorithm to encode a list of strings to a single string. The encoded string is then decoded back to the original list of strings.

Please implement encode and decode

Example 1:

Input: ["neet","code","love","you"]

Output:["neet","code","love","you"]
Example 2:

Input: ["we","say",":","yes"]

Output: ["we","say",":","yes"]
Constraints:

0 <= strs.length < 100
0 <= strs[i].length < 200
strs[i] contains only UTF-8 characters.
*/

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// Encode with len first, separated by '#', and then the word after
	// so we can easily identify which component is the nums and then slice
	// out the word using a start and end index
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(fmt.Sprintf("%d#%s", len(str), str))
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}

	var res []string

	i, j := 0, 0
	for j < len(encoded) {
		if encoded[j] == '#' {
			if currStrLen, err := strconv.Atoi(encoded[i:j]); err != nil {
				return nil
			} else {
				res = append(res, encoded[j+1:j+currStrLen+1])
				j += currStrLen + 1
				i = j
			}
		}
		j++
	}
	return res
}
