package binary_tree

import "math"

/*
Given a binary tree, return true if it is height-balanced and false otherwise.

A height-balanced binary tree is defined as a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

Example 1:

Input: root = [1,2,3,null,null,4]

Output: true
Example 2:

Input: root = [1,2,3,null,null,4,null,5]

Output: false
Example 3:

Input: root = []

Output: true
Constraints:

The number of nodes in the tree is in the range [0, 1000].
-1000 <= Node.val <= 1000
*/

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	balanced := true

	var getHeight func(*TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := getHeight(root.Left), getHeight(root.Right)
		if math.Abs(float64(left-right)) > 1 {
			balanced = false
		}
		return 1 + max(left, right)
	}
	getHeight(root)

	return balanced
}
