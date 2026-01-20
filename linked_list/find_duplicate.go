package linked_list

/*
You are given an array of integers nums containing n + 1 integers. Each integer in nums is in the range [1, n] inclusive.

Every integer appears exactly once, except for one integer which appears two or more times. Return the integer that appears more than once.

Example 1:

Input: nums = [1,2,3,2,2]

Output: 2
Example 2:

Input: nums = [1,2,3,4,4]

Output: 4
Follow-up: Can you solve the problem without modifying the array nums and using
O(1) extra space?

Constraints:

1 <= n <= 10000
nums.length == n + 1
1 <= nums[i] <= n
*/

func findDuplicate(nums []int) int {
	// since the numbers are constrained b/w 1 and n, we can use these numbers as indices to visit; if we visit the
	// same index again, then this means there is a duplicate. To do it w/o modifying the arr, we can treat this as a
	// linked list, which is the same as cycle detection.
	tortoise, hare := 0, 0
	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	// we have detected a cycle, but not necessarily the start of the cycle (which is the repeated number). Mathematically,
	// we have shown in the cycle detection problem that mathematically, the distance between the meeting point in the cycle
	// and the distance between the start of the arr to the start of the cycle is equivalent. LeetcodeReview/linked_list/cycle_detection.go.
	// so they just need to move one space each time until they meet.
	p := 0
	for {
		tortoise = nums[tortoise]
		p = nums[p]
		if p == tortoise {
			return tortoise
		}
	}
}
