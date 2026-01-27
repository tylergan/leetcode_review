package binary_tree

/*
Given the root of a binary tree, return its depth.

The depth of a binary tree is defined as the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:

Input: root = [1,2,3,null,null,4]

Output: 3
Example 2:

Input: root = []

Output: 0
Constraints:

0 <= The number of nodes in the tree <= 100.
-100 <= Node.val <= 100
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := maxDepth(root.Left)
	rightHeight := maxDepth(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

type Deque struct {
	q []*TreeNode
}

func NewDeque() Deque {
	return Deque{q: []*TreeNode{}}
}

func (d *Deque) Len() int { return len(d.q) }

func (d *Deque) Enqueue(node *TreeNode) {
	d.q = append(d.q, node)
}

func (d *Deque) Dequeue() *TreeNode {
	if len(d.q) == 0 {
		return nil
	}
	first := d.q[0]
	d.q = d.q[1:]
	return first
}

func maxDepthWithBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	deque := NewDeque()
	deque.Enqueue(root)
	level := 0
	for deque.Len() > 0 {
		sz := deque.Len()
		for i := 0; i < sz; i++ {
			currNode := deque.Dequeue()
			if currNode.Left != nil {
				deque.Enqueue(currNode.Left)
			}
			if currNode.Right != nil {
				deque.Enqueue(currNode.Right)
			}
		}
		level++
	}
	return level
}
