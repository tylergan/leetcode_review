package backtrack

/*
You are given an integer n. Return all well-formed parentheses strings that you can generate with n pairs of parentheses.

Example 1:

Input: n = 1

Output: ["()"]
Example 2:

Input: n = 3

Output: ["((()))","(()())","(())()","()(())","()()()"]
You may return the answer in any order.

Constraints:

1 <= n <= 7
*/

func generateParenthesis(n int) []string {
	res := []string{}

	var dfs func(string, int, int)
	dfs = func(s string, open, close int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}

		if open < n {
			dfs(s+"(", open+1, close)
		}

		if close < open {
			dfs(s+")", open, close+1)
		}
	}

	dfs("", 0, 0)

	return res
}
