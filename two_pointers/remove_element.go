package two_pointers

/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.

Example 1:

Input: nums = [3,2,2,3], val = 3

Output: k = 2, nums = [2,2,_,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).

Example 2:

Input: nums = [0,1,2,2,3,0,4,2], val = 2

Output: k = 5, nums = [0,1,3,0,4,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100
*/

// Treat j as the end of the still-valid region: shrink it past any trailing
// vals before swapping one in for a val found at i. j may never move (no val at
// the tail) or collapse straight onto i (array is all vals), so the i <= j and
// i < j guards keep single-element and no-swap cases from over-counting.
func removeElement(nums []int, val int) int {
	n := len(nums)
	j := n - 1
	i := 0

	for i <= j { // i <= j to account for single lengths and arrays where the array contains no val
		if nums[i] == val {
			for j >= i && nums[j] == val {
				j--
			}
			if i < j { // if i >= j, then that means no swap can be performed as we have completed all swaps
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		} else {
			i++
		}
	}

	return i
}
