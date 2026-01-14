package binary_search

import "math"

/*
You are given two integer arrays nums1 and nums2 of size m and n respectively, where each is sorted in ascending order. Return the median value among all elements of the two arrays.

Your solution must run in O(log(m+n)) time.

Example 1:

Input: nums1 = [1,2], nums2 = [3]

Output: 2.0
Explanation: Among [1, 2, 3] the median is 2.

Example 2:

Input: nums1 = [1,3], nums2 = [2,4]

Output: 2.5
Explanation: Among [1, 2, 3, 4] the median is (2 + 3) / 2 = 2.5.

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
-10^6 <= nums1[i], nums2[i] <= 10^6
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n2 < n1 {
		return findMedianSortedArrays(nums2, nums1) // conduct binary search on shorter arr since it CAN represent our entire left partition
	}

	n := n1 + n2
	m := n / 2

	l, r := 0, n1 // binary search on smaller arr
	for {
		i := (l + r) / 2 // one part of our left partition from the smaller arr
		j := m - i       // get the remaining elements for our left partition from the larger arr

		// Get our current last elem from first-part left partition and first elem from right partition from our smaller arr.
		aLeft := math.MinInt64
		if i > 0 { // if we go too far to left
			aLeft = nums1[i-1]
		}
		aRight := math.MaxInt64
		if i < n1 { // if we go too far to right
			aRight = nums1[i]
		}

		// Get our current last elem from rem. left partition and first elem from right partition from our larger arr
		bLeft := math.MinInt64
		if j > 0 {
			bLeft = nums2[j-1]
		}
		bRight := math.MaxInt64
		if j < n2 {
			bRight = nums2[j]
		}

		if aLeft <= bRight && bLeft <= aRight { // check that our maximal left partition elements from arr1 and arr2 are smaller than our minimal right partition elements
			if n%2 != 0 {
				return float64(min(aRight, bRight))
			} else {
				return float64(max(aLeft, bLeft)+min(aRight, bRight)) / 2.0
			}
		} else if aLeft > bRight { // if our smaller arr maximal left partition element is too large, look for smaller
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
