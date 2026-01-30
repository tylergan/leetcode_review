package binary_tree

/*
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) in the tree.

A binary search tree satisfies the following constraints:

The left subtree of every node contains only nodes with keys less than the node's key.
The right subtree of every node contains only nodes with keys greater than the node's key.
Both the left and right subtrees are also binary search trees.
Example 1:

Input: root = [2,1,3], k = 1

Output: 1
Example 2:

Input: root = [4,3,5,2,null], k = 4

Output: 5
Constraints:

1 <= k <= The number of nodes in the tree <= 1000.
0 <= Node.val <= 1000
*/

func kthSmallest(root *TreeNode, k int) int {
	ans := -1
	if root == nil {
		return ans
	}

	i := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		i++
		if i == k {
			ans = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func kthSmallestStack(root *TreeNode, k int) int {
	var stk []*TreeNode
	curr := root
	for curr != nil && len(stk) > 0 {
		for curr != nil {
			stk = append(stk, curr)
			curr = curr.Left
		}
		curr = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		k--
		if k == 0 {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}

func kthSmallestMorris(root *TreeNode, k int) int { // https://www.youtube.com/watch?v=wGXB9OWhPTg
	curr := root
	for curr != nil {
		if curr.Left == nil { // as per inorder, left --> root --> right
			k--
			if k == 0 {
				return curr.Val
			}
			curr = curr.Right
		} else {
			pred := curr.Left
			for pred.Right != nil && pred.Right != curr { // find the inorder pred. (though it is possible we have connected it before)
				pred = pred.Right
			}
			if pred.Right == nil { // this path has not been visited before, so we connect the inorder pred to the curr so that we can go back to curr later
				pred.Right = curr
				curr = curr.Left
			} else {
				// if we have traversed this path before to find the inorder pred. for the curr node, this suggests we
				//have visited the left side before, so visit the curr node and remove the link for the inorder pred.
				pred.Right = nil
				k--
				if k == 0 {
					return curr.Val
				}
				curr = curr.Right
			}
		}
	}
	return -1
}
