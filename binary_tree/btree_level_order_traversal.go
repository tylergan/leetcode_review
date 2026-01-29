package binary_tree

/*
Given a binary tree root, return the level order traversal of it as a nested list, where each sublist contains the values of nodes at a particular level in the tree, from left to right.

Example 1:

Input: root = [1,2,3,4,5,6,7]

Output: [[1],[2,3],[4,5,6,7]]
Example 2:

Input: root = [1]

Output: [[1]]
Example 3:

Input: root = []

Output: []
Constraints:

0 <= The number of nodes in the tree <= 1000.
-1000 <= Node.val <= 1000
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	dq := NewDeque()
	dq.Enqueue(root)
	for dq.Len() > 0 {
		length := dq.Len()
		subArr := make([]int, length)
		for i := 0; i < length; i++ {
			curr := dq.Dequeue()
			subArr[i] = curr.Val
			if curr.Left != nil {
				dq.Enqueue(curr.Left)
			}
			if curr.Right != nil {
				dq.Enqueue(curr.Right)
			}
		}
		res = append(res, subArr)
	}
	return res
}
