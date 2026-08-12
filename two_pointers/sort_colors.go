package two_pointers

/*
You are given an array nums consisting of n elements where each element is an integer representing a color:

0 represents red
1 represents white
2 represents blue

Your task is to sort the array in-place such that elements of the same color are grouped together and arranged in the order: red (0), white (1), and then blue (2).

You must not use any built-in sorting functions to solve this problem.

Example 1:

Input: nums = [1,0,1,2]

Output: [0,1,1,2]

Example 2:

Input: nums = [2,1,0]

Output: [0,1,2]

Constraints:

1 <= nums.length <= 300.
0 <= nums[i] <= 2.

Follow up: Could you come up with a one-pass algorithm using only constant extra space?
*/

// sortColors is the one-pass Dutch National Flag partition. It maintains:
//
//	[0, l)   -> all 0s (finalized)
//	[l, i)   -> all 1s (finalized)
//	[i, r]   -> unexamined
//	(r, n-1] -> all 2s (finalized)
//
// When nums[i] == 0 we swap it with nums[l]: the value coming back to i is from
// the [l, i) region (all 1s) or is a self-swap, so it is already validated and i
// advances. When nums[i] == 2 we swap with nums[r]: that value came from the
// unexamined region, so we must re-check position i (hence i--).
func sortColors(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	for i := l; i <= r; i++ {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		} else if nums[i] == 2 {
			nums[r], nums[i] = nums[i], nums[r]
			r--
			i-- // re-check the swapped-in value: it has not been examined yet
		}
	}
}

// sortColorsCounting is the two-pass counting sort. It uses O(1) extra space (a
// fixed-size frequency array) but reads the array twice. Intuitive, though it
// does not satisfy the one-pass follow-up.
func sortColorsCounting(nums []int) {
	var colorFreq [3]int
	for _, c := range nums {
		colorFreq[c]++
	}

	curr := 0
	for i := range nums {
		for curr < len(colorFreq) && colorFreq[curr] == 0 {
			curr++
		}
		nums[i] = curr
		colorFreq[curr]--
	}
}
