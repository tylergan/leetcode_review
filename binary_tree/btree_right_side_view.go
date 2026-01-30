package binary_tree

/*
You are given the root of a binary tree. Return only the values of the nodes that are visible from the right side of the tree, ordered from top to bottom.

Example 1:

Input: root = [1,2,3]

Output: [1,3]
Example 2:

Input: root = [1,2,3,4,5,6,7]

Output: [1,3,7]
*/

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int

	dq := NewDeque()
	dq.Enqueue(root)
	for dq.Len() > 0 {
		length := dq.Len()
		for i := 0; i < length; i++ {
			curr := dq.Dequeue()
			if i == length-1 {
				res = append(res, curr.Val)
			}
			if curr.Left != nil {
				dq.Enqueue(curr.Left)
			}
			if curr.Right != nil {
				dq.Enqueue(curr.Right)
			}
		}
	}
	return res
}

func rightSideViewDFS(root *TreeNode) []int {
	var res []int
	level := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, currLevel int) {
		if node == nil {
			return
		}

		if currLevel > level {
			res = append(res, node.Val)
			level = currLevel
		}
		dfs(node.Right, currLevel+1)
		dfs(node.Left, currLevel+1)
	}

	dfs(root, 1)
	return res
}
