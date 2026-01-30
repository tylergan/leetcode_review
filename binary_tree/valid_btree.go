package binary_tree

import (
	"math"
)

/*
Given the root of a binary tree, return true if it is a valid binary search tree, otherwise return false.

A valid binary search tree satisfies the following constraints:

The left subtree of every node contains only nodes with keys less than the node's key.
The right subtree of every node contains only nodes with keys greater than the node's key.
Both the left and right subtrees are also binary search trees.
Example 1:

Input: root = [2,1,3]

Output: true
Example 2:

Input: root = [1,2,3]

Output: false
Explanation: The root node's value is 1 but its left child's value is 2 which is greater than 1.

Constraints:

1 <= The number of nodes in the tree <= 1000.
-1000 <= Node.val <= 1000
*/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, lower int, upper int) bool {
		if node == nil {
			return true
		}
		if !(lower < node.Val && node.Val < upper) {
			return false
		}
		return dfs(node.Left, lower, node.Val) &&
			dfs(node.Right, node.Val, upper)
	}

	return dfs(root, math.MinInt32, math.MaxInt32)
}

func isValidBSTBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	type QueueItem struct {
		node  *TreeNode
		lower int
		upper int
	}

	q := []QueueItem{{root, math.MinInt32, math.MaxInt32}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		lower, upper, node := curr.lower, curr.upper, curr.node

		if !(lower < node.Val && node.Val < upper) {
			return false
		}
		if node.Left != nil {
			q = append(q, QueueItem{node.Left, lower, node.Val})
		}
		if node.Right != nil {
			q = append(q, QueueItem{node.Right, node.Val, upper})
		}
	}
	return true
}
