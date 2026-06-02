package graphs

/*
You are given a rectangular island heights where heights[r][c] represents the height above sea level of the cell at coordinate (r, c).

The islands borders the Pacific Ocean from the top and left sides, and borders the Atlantic Ocean from the bottom and right sides.

Water can flow in four directions (up, down, left, or right) from a cell to a neighboring cell with height equal or lower. Water can also flow into the ocean from cells adjacent to the ocean.

Find all cells where water can flow from that cell to both the Pacific and Atlantic oceans. Return it as a 2D list where each element is a list [r, c] representing the row and column of the cell. You may return the answer in any order.

Example 1:

Input: heights = [
  [4,2,7,3,4],
  [7,4,6,4,7],
  [6,3,5,3,6]
]

Output: [[0,2],[0,4],[1,0],[1,1],[1,2],[1,3],[1,4],[2,0]]
Example 2:

Input: heights = [[1],[1]]

Output: [[0,0],[1,0]]
Constraints:

1 <= heights.length, heights[r].length <= 100
0 <= heights[r][c] <= 1000
*/

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific := map[[2]int]bool{}
	atlantic := map[[2]int]bool{}

	var dfs func(int, int, map[[2]int]bool)
	dfs = func(i int, j int, ocean map[[2]int]bool) {
		ocean[[2]int{i, j}] = true

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			nxtRow, nxtCol := i+dir[0], j+dir[1]
			if nxtRow < 0 || nxtRow >= m || nxtCol < 0 || nxtCol >= n || heights[i][j] > heights[nxtRow][nxtCol] {
				continue
			}

			if _, ok := ocean[[2]int{nxtRow, nxtCol}]; ok {
				continue
			}

			dfs(nxtRow, nxtCol, ocean)
		}
	}

	// Start from left and right sides.
	for i := range m {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	// Start from top and bottom sides.
	for j := range n {
		dfs(0, j, pacific)
		dfs(m-1, j, atlantic)
	}

	// Find coords that are present in both the pacific and atlantic sets.
	var res [][]int
	for coords := range pacific {
		if _, ok := atlantic[coords]; ok {
			res = append(res, []int{coords[0], coords[1]})
		}
	}

	return res
}
