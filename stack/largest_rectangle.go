package stack

/*
You are given an array of integers heights where heights[i] represents the height of a bar. The width of each bar is 1.

Return the area of the largest rectangle that can be formed among the bars.

Note: This chart is known as a histogram.

Example 1:

Input: heights = [7,1,7,2,2,4]

Output: 8
Example 2:

Input: heights = [1,3,7]

Output: 7
Constraints:

1 <= heights.length <= 1000.
0 <= heights[i] <= 1000
*/

func largestRectangleArea(heights []int) int {
	maxArea := 0
	var stack []int // a mono increasing stack
	for i := 0; i <= len(heights); i++ {
		for len(stack) > 0 && (i == len(heights) || heights[stack[len(stack)-1]] >= heights[i]) { // one or more bars these cannot be extended beyond our current rightBound (heights[i]), or we need to deplete the stack as we are at the end with still bars leftover.
			// calculate the rectangle area for our current height
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i          // this is our rightBoundIdx.
			if len(stack) > 0 { // Check if there is a leftBoundIdx, otherwise it has no left bound.
				leftBoundIdx := stack[len(stack)-1]
				width = i - leftBoundIdx - 1 // -1 to exclude the left boundary from the width, giving us the rectangle for just the bar on the top of the stack
			}
			maxArea = max(maxArea, width*height)
		}
		if i < len(heights) {
			stack = append(stack, i)
		}
	}
	return maxArea
}
