package binary_tree

import "math"

/*
Given the root of a non-empty binary tree, return the maximum path sum of any non-empty path.

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge connecting them. A node can not appear in the sequence more than once. The path does not necessarily need to include the root.

The path sum of a path is the sum of the node's values in the path.

Example 1:



Input: root = [1,2,3]

Output: 6
Explanation: The path is 2 -> 1 -> 3 with a sum of 2 + 1 + 3 = 6.

Example 2:



Input: root = [-15,10,20,null,null,15,5,-5]

Output: 40
Explanation: The path is 15 -> 20 -> 5 with a sum of 15 + 20 + 5 = 40.
*/

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs(node.Left), dfs(node.Right)

		// for any given node, we compare res to a potential path that includes the path from the left and right children
		currVal := node.Val
		if left > 0 {
			currVal += left
		}
		if right > 0 {
			currVal += right
		}
		res = max(res, currVal)

		// a path cannot branch, so the current node should only extend from the best path (left or right)
		// and return the best path to the parent
		if left < 0 && right < 0 {
			return node.Val
		}
		if left > right {
			return node.Val + left
		}
		return node.Val + right
	}

	dfs(root)
	return res
}

// //  A cleaner solution
// func maxPathSum(root *TreeNode) int {
// 	res := math.MinInt32

// 	var dfs func(*TreeNode) int
// 	dfs = func(node *TreeNode) int {
// 		if node == nil {
// 			return 0
// 		}
// 		left := max(dfs(node.Left), 0)
// 		right := max(dfs(node.Right), 0)

// 		// for any given node, we compare res to a potential path that includes the path from the left and right children
// 		currVal := node.Val + left + right
// 		res = max(res, currVal)

// 		// a path cannot branch, so the current node should only extend from the best path (left or right)
// 		// and return the best path to the parent
// 		return node.Val + max(left, right)
// 	}

// 	dfs(root)
// 	return res
// }
