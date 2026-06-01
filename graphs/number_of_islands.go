package graphs

/*
Given a 2D grid grid where '1' represents land and '0' represents water, count and return the number of islands.

An island is formed by connecting adjacent lands horizontally or vertically and is surrounded by water. You may assume water is surrounding the grid (i.e., all the edges are water).

Example 1:

Input: grid = [
    ["0","1","1","1","0"],
    ["0","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
  ]
Output: 1
Example 2:

Input: grid = [
    ["1","1","0","0","1"],
    ["1","1","0","0","1"],
    ["0","0","1","0","0"],
    ["0","0","0","1","1"]
  ]
Output: 4
Constraints:

1 <= grid.length, grid[i].length <= 100
grid[i][j] is '0' or '1'.
*/

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	count := 0

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		if grid[i][j] == byte('0') {
			return
		}

		grid[i][j] = byte('0')

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}

	for i := range m {
		for j := range n {
			if grid[i][j] == byte('0') {
				continue
			}

			count++
			dfs(i, j)
		}
	}

	return count
}
