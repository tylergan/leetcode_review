package backtrack

import "slices"

/*
Given an array nums of unique integers, return all possible subsets of nums.

The solution set must not contain duplicate subsets. You may return the solution in any order.

Example 1:

Input: nums = [1,2,3]

Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [7]

Output: [[],[7]]
Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/

func subsets(nums []int) [][]int {
	res := [][]int{{}}
	curr := []int{}

	var dfs func([]int)
	dfs = func(nums []int) {
		if len(curr) > 0 {
			res = append(res, slices.Clone(curr))
		}
		for i, n := range nums {
			curr = append(curr, n)
			if i < len(nums) {
				dfs(nums[i+1:])
			} else {
				dfs([]int{})
			}
			curr = curr[:len(curr)-1]
		}
	}

	dfs(nums)
	return res
}
