package backtrack

import (
	"slices"
	"sort"
)

/*
You are given an array of integers candidates, which may contain duplicates, and a target integer target. Your task is to return a list of all unique combinations of candidates where the chosen numbers sum to target.

Each element from candidates may be chosen at most once within a combination. The solution set must not contain duplicate combinations.

You may return the combinations in any order and the order of the numbers in each combination can be in any order.

Example 1:

Input: candidates = [9,2,2,4,6,1,5], target = 8

Output: [
  [1,2,5],
  [2,2,4],
  [2,6]
]
Example 2:

Input: candidates = [1,2,3,4,5], target = 7

Output: [
  [1,2,4],
  [2,5],
  [3,4]
]
Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(int, []int, int)
	dfs = func(i int, curr []int, currTotal int) {
		if currTotal == target {
			res = append(res, slices.Clone(curr))
			return
		}
		for j := i; j < len(candidates); j++ {
			if currTotal+candidates[j] > target {
				return
			}
			curr = append(curr, candidates[j])
			dfs(j+1, curr, currTotal+candidates[j])
			curr = curr[:len(curr)-1]
			for j+1 < len(candidates) && candidates[j] == candidates[j+1] {
				j++
			}
		}
	}
	dfs(0, []int{}, 0)
	return res
}
