package binary_tree

import "math"

/*
Within a binary tree, a node x is considered good if the path from the root of the tree to the node x contains no nodes with a value greater than the value of node x

Given the root of a binary tree root, return the number of good nodes within the tree.

Example 1:

Input: root = [2,1,1,3,null,1,5]

Output: 3

Example 2:

Input: root = [1,2,-1,3,4]

Output: 4
Constraints:

1 <= number of nodes in the tree <= 100
-100 <= Node.val <= 100
*/

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, largestVal int) {
		if node == nil {
			return
		}
		if node.Val >= largestVal {
			count++
		}
		dfs(node.Left, max(node.Val, largestVal))
		dfs(node.Right, max(node.Val, largestVal))
	}
	dfs(root, root.Val)
	return count
}

func goodNodesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type NodeData struct {
		node   *TreeNode
		maxVal int
	}

	count := 0
	q := []NodeData{{root, math.MinInt32}}

	for len(q) > 0 {
		currPath := q[0]
		q = q[1:]

		node := currPath.node
		maxVal := currPath.maxVal
		if node.Val >= maxVal {
			count++
		}
		if node.Left != nil {
			q = append(q, NodeData{node.Left, max(maxVal, node.Val)})
		}
		if node.Right != nil {
			q = append(q, NodeData{node.Right, max(maxVal, node.Val)})
		}
	}

	return count
}
