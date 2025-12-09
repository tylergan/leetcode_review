package two_pointers

/*
You are given an integer array heights where heights[i] represents the height of the ith bar.

You may choose any two bars to form a container. Return the maximum amount of water a container can store.

Example 1:

Input: height = [1,7,2,5,4,7,3,6]

Output: 36
Example 2:

Input: height = [2,2,2]

Output: 4
Constraints:

2 <= height.length <= 1000
0 <= height[i] <= 1000
*/

func maxArea(heights []int) int {
	res := 0
	i, j := 0, len(heights)-1
	for i < j {
		res = max(res, min(heights[i], heights[j])*(j-i))
		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return res
}
