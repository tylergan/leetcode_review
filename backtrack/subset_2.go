package backtrack

import (
	"slices"
	"sort"
)

/*
You are given an array nums of integers, which may contain duplicates. Return all possible subsets.

The solution must not contain duplicate subsets. You may return the solution in any order.

Example 1:

Input: nums = [1,2,1]

Output: [[],[1],[1,2],[1,1],[1,2,1],[2]]
Example 2:

Input: nums = [7,7]

Output: [[],[7], [7,7]]
Constraints:

1 <= nums.length <= 11
-20 <= nums[i] <= 20
*/

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	curr := []int{}

	sort.Ints(nums)

	var dfs func(int)
	dfs = func(i int) {
		res = append(res, slices.Clone(curr))

		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}

			curr = append(curr, nums[j])
			dfs(j + 1)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(0)

	return res
}
