package backtrack

import "slices"

/*
Given an array nums of unique integers, return all the possible permutations. You may return the answer in any order.

Example 1:

Input: nums = [1,2,3]

Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [7]

Output: [[7]]
Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
*/

func permute(nums []int) [][]int {
	res := [][]int{}
	curr := []int{}
	seen := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(curr) == len(nums) {
			res = append(res, slices.Clone(curr))
			return
		}
		for i, n := range nums {
			if seen[i] {
				continue
			}
			curr = append(curr, n)
			seen[i] = true
			dfs()
			curr = curr[:len(curr)-1]
			seen[i] = false
		}
	}
	dfs()
	return res
}
