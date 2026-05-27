package backtrack

import "strings"

/*
The n-queens puzzle is the problem of placing n queens on an n x n chessboard so that no two queens can attack each other.

A queen in a chessboard can attack horizontally, vertically, and diagonally.

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a unique board layout where the queen pieces are placed. 'Q' indicates a queen and '.' indicates an empty space.

You may return the answer in any order.

Example 1:

Input: n = 4

Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
Explanation: There are two different solutions to the 4-queens puzzle.

Example 2:

Input: n = 1

Output: [["Q"]]
Constraints:

1 <= n <= 8
*/

func solveNQueens(n int) [][]string {
	var res [][]string

	curr := createMatrix(n)

	var dfs func(int)
	dfs = func(r int) {
		if r == n {
			res = append(res, toString(curr))
			return
		}

		for c := range n {
			if isSafe(r, c, curr) {
				curr[r][c] = 'Q'
				dfs(r + 1)
				curr[r][c] = '.'
			}
		}
	}

	dfs(0)

	return res
}

func createMatrix(n int) [][]rune {
	var curr [][]rune
	for range n {
		var row strings.Builder
		for range n {
			row.WriteString(".")
		}
		curr = append(curr, []rune(row.String()))
	}
	return curr
}

func toString(matrix [][]rune) []string {
	res := make([]string, len(matrix))
	for i, row := range matrix {
		res[i] = string(row)
	}
	return res
}

func isSafe(r int, c int, matrix [][]rune) bool {
	// Since all prior queen placements are on rows above the current row,
	// there is no need to check the current row horizontally.
	n := len(matrix)

	// Check vertical.
	for i := range r {
		if matrix[i][c] == 'Q' {
			return false
		}
	}

	checkSlashParts := func(rowStep, colStep int) bool {
		for step := 1; step < n; step++ {
			rIdx := r + step*rowStep
			cIdx := c + step*colStep
			if rIdx < 0 || rIdx >= n || cIdx < 0 || cIdx >= n {
				break
			}

			if matrix[rIdx][cIdx] == 'Q' {
				return false
			}
		}

		return true
	}

	// Check forward slash.
	forwardBot, forwardUp := checkSlashParts(1, 1), checkSlashParts(-1, -1)
	if !forwardBot || !forwardUp {
		return false
	}

	// Check backward slash.
	backwardUp, backwardDown := checkSlashParts(-1, 1), checkSlashParts(1, -1)
	if !backwardUp || !backwardDown {
		return false
	}

	return true
}
