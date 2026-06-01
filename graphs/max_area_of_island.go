package graphs

/*
You are given a matrix grid where grid[i] is either a 0 (representing water) or 1 (representing land).

An island is defined as a group of 1's connected horizontally or vertically. You may assume all four edges of the grid are surrounded by water.

The area of an island is defined as the number of cells within the island.

Return the maximum area of an island in grid. If no island exists, return 0.

Example 1:

Input: grid = [
  [0,1,1,0,1],
  [1,0,1,0,1],
  [0,1,1,0,1],
  [0,1,0,0,1]
]

Output: 6
Explanation: 1's cannot be connected diagonally, so the maximum area of the island is 6.

Constraints:

1 <= grid.length, grid[i].length <= 50
*/

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	maxArea := 0
	currArea := 0

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == 0 {
			return
		}

		currArea++
		grid[i][j] = 0

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				continue
			}

			dfs(i, j)

			maxArea = max(maxArea, currArea)
			currArea = 0
		}
	}

	return maxArea
}
