package graphs

/*
You are given a 2-D matrix grid. Each cell can have one of three possible values:

0 representing an empty cell
1 representing a fresh fruit
2 representing a rotten fruit
Every minute, if a fresh fruit is horizontally or vertically adjacent to a rotten fruit, then the fresh fruit also becomes rotten.

Return the minimum number of minutes that must elapse until there are zero fresh fruits remaining. If this state is impossible within the grid, return -1.

Example 1:

Input: grid = [[1,1,0],[0,1,1],[0,1,2]]

Output: 4
Example 2:

Input: grid = [[1,0,1],[0,2,0],[1,0,1]]

Output: -1
Constraints:

1 <= grid.length, grid[i].length <= 10
*/

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	totFruits := 0 // Keep track of total fruits in grid.

	var q [][]int
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				continue
			}

			if grid[i][j] == 2 {
				q = append(q, []int{i, j, 0}) // Rotting fruit location, starting at t=0.
			}

			totFruits++
		}
	}

	cntFruits := 0 // Number of fruits reached by the BFS flood from starting points.
	maxTime := 0
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		curRow, curCol, curTime := cur[0], cur[1], cur[2]
		cntFruits++

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			nxtTime := curTime + 1
			nxtRow, nxtCol := curRow+dir[0], curCol+dir[1]

			if nxtRow < 0 || nxtRow >= m || nxtCol < 0 || nxtCol >= n || grid[nxtRow][nxtCol] != 1 {
				continue
			}

			grid[nxtRow][nxtCol] = 2
			q = append(q, []int{nxtRow, nxtCol, nxtTime})
			maxTime = max(maxTime, nxtTime)
		}
	}

	if totFruits != cntFruits {
		return -1
	}
	return maxTime
}
