package graphs

/*
You are given a m x n 2D grid initialized with these three possible values:

-1 - A water cell that can not be traversed.
0 - A treasure chest.
INF - A land cell that can be traversed. We use the integer 2^31 - 1 = 2147483647 to represent INF.

Fill each land cell with the distance to its nearest treasure chest. If a land cell cannot reach a treasure chest then the value should remain INF.

Assume the grid can only be traversed up, down, left, or right.

Modify the grid in-place.

Example 1:

Input: [
  [2147483647,-1,0,2147483647],
  [2147483647,2147483647,2147483647,-1],
  [2147483647,-1,2147483647,-1],
  [0,-1,2147483647,2147483647]
]

Output: [
  [3,-1,0,1],
  [2,2,1,-1],
  [1,-1,2,-1],
  [0,-1,3,4]
]
Example 2:

Input: [
  [0,-1],
  [2147483647,2147483647]
]

Output: [
  [0,-1],
  [1,2]
]
*/

func islandsAndTreasure(grid [][]int) {
	INF := 2147483647

	m, n := len(grid), len(grid[0])

	q := [][]int{}

	// Find all treasure chests.
	for i := range m {
		for j := range n {
			if grid[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}

	// Build distances for each neighbour cell from treasure chest.
	for len(q) > 0 {
		cur := q[0]
		curRow, curCol := cur[0], cur[1]

		q = q[1:]

		dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range dirs {
			nxtRow, nxtCol := curRow+dir[0], curCol+dir[1]
			if nxtRow < 0 || nxtRow >= m || nxtCol < 0 || nxtCol >= n || grid[nxtRow][nxtCol] != INF {
				continue
			}

			q = append(q, []int{nxtRow, nxtCol})
			grid[nxtRow][nxtCol] = grid[curRow][curCol] + 1
		}
	}
}
