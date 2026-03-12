package binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Implement an algorithm to serialize and deserialize a binary tree.

Serialization is the process of converting an in-memory structure into a sequence of bits so that it can be stored or sent across a network to be reconstructed later in another computer environment.

You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure. There is no additional restriction on how your serialization/deserialization algorithm should work.

Note: The input/output format in the examples is the same as how NeetCode serializes a binary tree. You do not necessarily need to follow this format.

Example 1:

Input: root = [1,2,3,null,null,4,5]

Output: [1,2,3,null,null,4,5]
Example 2:

Input: root = []

Output: []
Constraints:

0 <= The number of nodes in the tree <= 1000.
-1000 <= Node.val <= 1000
*/

type Codec struct{}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "$"
	}

	var res strings.Builder
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res.WriteString("$#")
			continue
		}
		res.WriteString(fmt.Sprintf("%d#", node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "$" {
		return nil
	}

	parts := strings.Split(data, "#")
	val, _ := strconv.Atoi(parts[0]) // the root should not be nil because serialise will return "$"
	root := &TreeNode{Val: val}
	queue := []*TreeNode{root}
	idx := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if idx < len(parts) {
			if parts[idx] != "$" {
				val, _ := strconv.Atoi(parts[idx])
				node.Left = &TreeNode{Val: val}
				queue = append(queue, node.Left)
			}
			idx++
		}
		if idx < len(parts) {
			if parts[idx] != "$" {
				val, _ := strconv.Atoi(parts[idx])
				node.Right = &TreeNode{Val: val}
				queue = append(queue, node.Right)
			}
			idx++
		}
	}
	return root
}
