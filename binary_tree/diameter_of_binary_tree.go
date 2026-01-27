package binary_tree

/*
The diameter of a binary tree is defined as the length of the longest path between any two nodes within the tree. The path does not necessarily have to pass through the root.

The length of a path between two nodes in a binary tree is the number of edges between the nodes. Note that the path can not include the same node twice.

Given the root of a binary tree root, return the diameter of the tree.

Example 1:

Input: root = [1,null,2,3,4,5]

Output: 3
Explanation: 3 is the length of the path [1,2,3,5] or [5,3,2,4].

Example 2:

Input: root = [1,2,3]

Output: 2
Constraints:

1 <= number of nodes in the tree <= 100
-100 <= Node.val <= 100
*/

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// True diameter of a tree must be entirely inside one subtree and never pass through the root of the whole tree.
	// If we only returned leftHeight + rightHeight, we are simply returning the diameter of paths that go through
	// the current root. However, we miss cases where the longest path is completely contained in the left or right
	// subtree. So, we need to keep track of the largest diameter we have seen so far, calculating the diameter of each
	// subtree.
	var getHeight func(*TreeNode) int
	getHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getHeight(root.Left)
		right := getHeight(root.Right)
		maxDiameter = max(maxDiameter, left+right)
		return 1 + max(left, right)
	}

	getHeight(root)
	return maxDiameter
}
