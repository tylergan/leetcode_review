package binary_tree

/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:

Input: root = [1,2,3,4,5], subRoot = [2,4,5]

Output: true
Example 2:

Input: root = [1,2,3,4,5,null,null,6], subRoot = [2,4,5]

Output: false
Constraints:

1 <= The number of nodes in both trees <= 100.
-100 <= root.val, subRoot.val <= 100
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var sameBinaryTree func(*TreeNode, *TreeNode) bool
	sameBinaryTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		return sameBinaryTree(p.Left, q.Left) && sameBinaryTree(p.Right, q.Right)
	}
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	} else if sameBinaryTree(root, subRoot) { // if sameBinaryTree fails, possible that there still exists a subtree e.g., root = [1, 1], sub-root = [1]
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
