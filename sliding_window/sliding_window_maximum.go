package sliding_window

/*
You are given an array of integers nums and an integer k. There is a sliding window of size k that starts at the left edge of the array. The window slides one position to the right until it reaches the right edge of the array.

Return a list that contains the maximum element in the window at each step.

Example 1:

Input: nums = [1,2,1,0,4,2,6], k = 3

Output: [2,2,4,4,6]

Explanation:
Window position            Max
---------------           -----
[1  2  1] 0  4  2  6        2
 1 [2  1  0] 4  2  6        2
 1  2 [1  0  4] 2  6        4
 1  2  1 [0  4  2] 6        4
 1  2  1  0 [4  2  6]       6
Constraints:

1 <= nums.length <= 1000
-10,000 <= nums[i] <= 10,000
1 <= k <= nums.length
*/

func maxSlidingWindow(nums []int, k int) []int {
	var output []int
	var deque []int
	i, j := 0, 0

	for j < len(nums) {
		// maintain the largest element at the end of the deque
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[j] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, j)

		// remove the older element from the start of the deque
		if i > deque[0] {
			deque = deque[1:]
		}

		// maintain window of size k
		if (j + 1) >= k {
			output = append(output, nums[deque[0]])
			i += 1
		}
		j += 1
	}

	return output
}
