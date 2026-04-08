package backtrack

/*
Given a 2-D grid of characters board and a string word, return true if the word is present in the grid, otherwise return false.

For the word to be present it must be possible to form it with a path in the board with horizontally or vertically neighboring cells. The same cell may not be used more than once in a word.

Example 1:



Input:
board = [
  ["A","B","C","D"],
  ["S","A","A","T"],
  ["A","C","A","E"]
],
word = "CAT"

Output: true
Example 2:



Input:
board = [
  ["A","B","C","D"],
  ["S","A","A","T"],
  ["A","C","A","E"]
],
word = "BAT"

Output: false

Approach: Backtracking with in-place visited marking
- Try starting the search from every cell that matches word[0].
- From each cell, DFS in all 4 directions (up, down, left, right), matching word[k] at each step.
- To avoid revisiting a cell in the same path, temporarily overwrite it with '#' before
  recursing, then restore it after (backtrack). This replaces the need for a separate visited set.
- Base cases:
    - k == len(word): all characters matched, return true.
    - Out of bounds or board[i][j] != word[k]: this path is invalid, return false.

Time: O(m * n * 4^L) where m*n is the board size and L is the word length.
      Each cell can branch into 4 directions, and we try every cell as a starting point.
Space: O(L) for the recursion stack (no extra visited matrix).
*/

func exist(board [][]byte, word string) bool {
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if k == len(word) {
			return true
		}

		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
			return false
		}

		ch := board[i][j]
		board[i][j] = '#'

		for _, dir := range dirs {
			di, dj := dir[0], dir[1]
			if dfs(i+di, j+dj, k+1) {
				board[i][j] = ch
				return true
			}
		}

		board[i][j] = ch
		return false
	}

	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
