package binary_search

/*
You are given an m x n 2-D integer array matrix and an integer target.

Each row in matrix is sorted in non-decreasing order.
The first integer of every row is greater than the last integer of the previous row.
Return true if target exists within matrix or false otherwise.

Can you write a solution that runs in O(log(m * n)) time?

Example 1:

Input: matrix = [
	[1,2,4,8],
	[10,11,12,13],
	[14,20,30,40]
], target = 10

Output: true
Example 2:

Input: matrix = [
	[1,2,4,8],
	[10,11,12,13],
	[14,20,30,40]
], target = 15

Output: false

Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-10000 <= matrix[i][j], target <= 10000
*/

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		if !(row[0] <= target && target <= row[len(row)-1]) {
			continue
		}

		searchRow := func() bool {
			l, r := 0, len(row)-1
			for l <= r {
				mid := (l + r) / 2
				if row[mid] == target {
					return true
				} else if row[mid] < target {
					l = mid + 1
				} else {
					r = mid - 1
				}
			}
			return false
		}

		return searchRow()
	}

	return false
}
