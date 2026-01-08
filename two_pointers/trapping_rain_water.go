package two_pointers

/*
You are given an array of non-negative integers height which represent an elevation map. Each value height[i] represents the height of a bar, which has a width of 1.

Return the maximum area of water that can be trapped between the bars.

Example 1:

Input: height = [0,2,0,3,1,0,1,3,2,1]

Output: 9
Constraints:

1 <= height.length <= 1000
0 <= height[i] <= 1000
*/

func trapEasierToUnderstand(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	// find the max left heights for each position
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// find the max right heights for each position
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// amount of water we can trap at some position 'i' is simply min(L, R) - height[i]
	res := 0
	for i := 0; i < n; i++ {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

// although this approach is not anymore efficient than the prior approach, this approach is good to see since this
// technique is similar to the one used for the stack-based question "maximum rectangle area"
func trapStackApproach(height []int) int {
	if len(height) == 0 {
		return 0
	}
	res := 0
	var stack []int // monotonic decreasing stack
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			// we see a wall that is taller than the bar at the top of the stack, we have found a rightBound
			rightBoundHeight := height[i]

			bottomIdx := stack[len(stack)-1] // this is the bottom between our two walls
			bottomHeight := height[bottomIdx]
			stack = stack[:len(stack)-1]

			// check for our left bound
			if len(stack) > 0 {
				leftBoundIdx := stack[len(stack)-1]
				leftBoundHeight := height[leftBoundIdx]
				h := min(leftBoundHeight, rightBoundHeight) - bottomHeight
				w := i - leftBoundIdx - 1
				res += h * w
			}
		}
		stack = append(stack, i)
	}
	return res
}

// trap is the more optimised approach, but the previous approach makes it easier for us to understand what we are doing
func trap(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]
	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			res += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			res += rightMax - height[r]
		}
	}
	return res
}
