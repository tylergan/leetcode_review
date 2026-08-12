package arrays

import "slices"

/*
You are given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.

Example 1:

Input: nums = [10,9,1,1,1,2,3,1]

Output: [1,1,1,1,2,3,9,10]

Example 2:

Input: nums = [5,10,2,1,3]

Output: [1,2,3,5,10]

Constraints:

1 <= nums.length <= 50,000.
-50,000 <= nums[i] <= 50,000
*/

// mergeSort sorts nums in place. The recursive calls' return values are
// discarded on purpose: the merge writes back into the slice it was handed, so
// sorting "left" and "right" happens in place. The halves MUST be cloned first
// because Go slices share the same backing array — without the clones, writing
// into nums during the merge would simultaneously overwrite the left/right data
// still being read.
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := slices.Clone(nums[:mid])
	right := slices.Clone(nums[mid:])

	mergeSort(left)
	mergeSort(right)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}

	return nums
}

func sortArray(nums []int) []int {
	return mergeSort(nums)
}
