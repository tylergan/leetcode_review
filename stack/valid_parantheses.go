package stack

/*
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.

Example 1:

Input: s = "[]"

Output: true
Example 2:

Input: s = "([{}])"

Output: true
Example 3:

Input: s = "[(])"

Output: false
Explanation: The brackets are not closed in the correct order.

Constraints:

1 <= s.length <= 1000
*/

func isValid(s string) bool {
	endPair := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var stack []string
	for _, c := range s {
		if match, ok := endPair[string(c)]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != match {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(c))
		}
	}
	return len(stack) == 0
}
