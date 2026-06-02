package graphs

/*
You are given a 2-D matrix board containing 'X' and 'O' characters.

If a continuous, four-directionally connected group of 'O's is surrounded by 'X's, it is considered to be surrounded.

Change all surrounded regions of 'O's to 'X's and do so in-place by modifying the input board.

Example 1:

Input: board = [
  ["X","X","X","X"],
  ["X","O","O","X"],
  ["X","O","O","X"],
  ["X","X","X","O"]
]

Output: [
  ["X","X","X","X"],
  ["X","X","X","X"],
  ["X","X","X","X"],
  ["X","X","X","O"]
]
Explanation: Note that regions that are on the border are not considered surrounded regions.

Constraints:

1 <= board.length, board[i].length <= 200
board[i][j] is 'X' or 'O'.
*/

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(int, int, byte, byte)
	dfs = func(i int, j int, target byte, replace byte) {
		board[i][j] = replace

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			nxtRow, nxtCol := i+dir[0], j+dir[1]
			if nxtRow < 0 || nxtRow >= m || nxtCol < 0 || nxtCol >= n || board[nxtRow][nxtCol] != target {
				continue
			}

			dfs(nxtRow, nxtCol, target, replace)
		}
	}

	// Mark components that have an edge to the outer parts of the board as safe.
	markSafe := func(target, replace byte) {
		for i := range m { // Left and right sides.
			if board[i][0] == target {
				dfs(i, 0, target, replace)
			}
			if board[i][n-1] == target {
				dfs(i, n-1, target, replace)
			}
		}

		for j := range n { // Top and bottom sides.
			if board[0][j] == target {
				dfs(0, j, target, replace)
			}
			if board[m-1][j] == target {
				dfs(m-1, j, target, replace)
			}
		}
	}

	changeFrom := func(target, replace byte) {
		for i := range m {
			for j := range n {
				if board[i][j] != target {
					continue
				}

				dfs(i, j, target, replace)
			}
		}
	}

	markSafe(byte('O'), byte('T'))
	changeFrom(byte('O'), byte('X'))
	changeFrom(byte('T'), byte('O'))
}
