package arrays

/*
Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.

You may return the output in any order.

Example 1:

Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
Example 2:

Input: nums = [7,7], k = 1

Output: [7]
Constraints:

1 <= nums.length <= 10^4.
-1000 <= nums[i] <= 1000
1 <= k <= number of distinct elements in nums.
*/

func topKFrequent(nums []int, k int) []int {
	// Get frequencies of nums and group them
	freqs := map[int]int{}
	for _, n := range nums {
		freqs[n]++
	}

	groupFreqs := map[int][]int{}
	for n, freq := range freqs {
		groupFreqs[freq] = append(groupFreqs[freq], n)
	}

	// Largest possible frequency is size of array. Check from n...0 to add group freq numbers to res.
	var res []int
	for i := len(nums); i >= 0 && len(res) < k; i-- {
		if group, ok := groupFreqs[i]; ok {
			res = append(res, group[:min(len(group), k-len(res))]...)
		}
	}

	return res
}
