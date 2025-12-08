package arrays

import "fmt"

/*
You are given a 9 x 9 Sudoku board. A Sudoku board is valid if the following rules are followed:

Each row must contain the digits 1-9 without duplicates.
Each column must contain the digits 1-9 without duplicates.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
Return true if the Sudoku board is valid, otherwise return false

Note: A board does not need to be full or be solvable to be valid.

Example 1:



Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","8",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: true
Example 2:

Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","1",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: false
Explanation: There are two 1's in the top-left 3x3 sub-box.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudoku(board [][]byte) bool {
	seen := map[string]map[byte]bool{}

	for i := 0; i < len(board); i++ {
		row := fmt.Sprintf("row-%d", i)
		if _, ok := seen[row]; !ok {
			seen[row] = map[byte]bool{}
		}

		rowSet, _ := seen[row]
		for j := 0; j < len(board); j++ {
			currCell := board[i][j]
			if currCell == '.' {
				continue
			}

			// Check if in row already
			if _, ok := rowSet[currCell]; ok {
				return false
			}
			rowSet[currCell] = true
			seen[row] = rowSet

			// Check in col already
			col := fmt.Sprintf("col-%d", j)
			if _, ok := seen[col]; !ok {
				seen[col] = map[byte]bool{}
			}
			colSet, _ := seen[col]
			if _, ok := colSet[currCell]; ok {
				return false
			}
			colSet[currCell] = true
			seen[col] = colSet

			// Check if in 3x3 block already
			blockNumber := int(i/3)*3 + j/3 // each time we go down, each row contains 3 blocks that we skip, so multiply row by 3
			currBlock := fmt.Sprintf("block-%d", blockNumber)
			if _, ok := seen[currBlock]; !ok {
				seen[currBlock] = map[byte]bool{}
			}
			blockSet := seen[currBlock]
			if _, ok := blockSet[currCell]; ok {
				return false
			}
			blockSet[currCell] = true
			seen[currBlock] = blockSet
		}
	}

	return true
}
