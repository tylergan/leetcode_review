package arrays

/*
Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].

Each product is guaranteed to fit in a 32-bit integer.

Follow-up: Could you solve it in O(n) time without using the division operation?

Example 1:

Input: nums = [1,2,4,6]

Output: [48,24,12,8]
Example 2:

Input: nums = [-1,0,1,2,3]

Output: [0,-6,0,0,0]
Constraints:

2 <= nums.length <= 1000
-20 <= nums[i] <= 20
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)

	// Multiply all numbers from left
	prefix := make([]int, n)
	prefix[0] = nums[0]
	for i := 1; i < n; i++ {
		prefix[i] = nums[i] * prefix[i-1]
	}

	// Multiply all numbers from right
	suffix := make([]int, n)
	suffix[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = nums[i] * suffix[i+1]
	}

	// Product excluding itself for curr position should be product
	// of left (before curr pos) and product of right (after curr pos)
	res := make([]int, n)
	res[0], res[n-1] = suffix[1], prefix[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = prefix[i-1] * suffix[i+1]
	}

	return res
}
