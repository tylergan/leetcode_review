package binary_search

import (
	"math"
	"slices"
)

/*
You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.

Return the minimum integer k such that you can eat all the bananas within h hours.

Example 1:

Input: piles = [1,4,3,2], h = 9

Output: 2
Explanation: With an eating rate of 2, you can eat the bananas in 6 hours. With an eating rate of 1, you would need 10 hours to eat all the bananas (which exceeds h=9), thus the minimum eating rate is 2.

Example 2:

Input: piles = [25,10,23,4], h = 4

Output: 25
Constraints:

1 <= piles.length <= 1,000
piles.length <= h <= 1,000,000
1 <= piles[i] <= 1,000,000,000
*/

func minEatingSpeed(piles []int, h int) int {
	largestPile := slices.Max(piles)
	l, r := 1, largestPile
	for l <= r {
		k := (l + r) / 2

		totalTime := 0
		for _, pile := range piles {
			totalTime += int(math.Ceil(float64(pile) / float64(k)))
		}

		if totalTime <= h {
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return l
}

/*
  Define P(k) = (totalTime(k) <= h).
  This predicate is monotonic: if P(k) is true, then P(k+1) is also true (faster speed never hurts). That means there’s a boundary where it flips from false → true. Your loop:

  - If P(mid) is true, you move high = mid - 1 to search for a smaller valid k.
  - If P(mid) is false, you move low = mid + 1 to search for a larger k.

  When the loop ends, low is the smallest k such that P(k) is true (the first valid speed). That’s why it “converges onto low.”

  It’s the same idea as “first index” binary search—here the “index” is the speed k.

  To find the last index that satisfies the predicate (i.e., the greatest index with P(index)=true), you flip the bias:

  - If P(mid) is true, move right: low = mid + 1
  - If P(mid) is false, move left: high = mid - 1
  - Return high at the end
*/
