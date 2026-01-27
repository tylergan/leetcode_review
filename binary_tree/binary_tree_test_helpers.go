package binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intPtr(v int) *int { return &v }

func buildTree(level []*int) *TreeNode {
	if len(level) == 0 || level[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *level[0]}
	queue := []*TreeNode{root}
	idx := 1
	for idx < len(level) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if idx < len(level) {
			if level[idx] != nil {
				node.Left = &TreeNode{Val: *level[idx]}
				queue = append(queue, node.Left)
			}
			idx++
		}
		if idx < len(level) {
			if level[idx] != nil {
				node.Right = &TreeNode{Val: *level[idx]}
				queue = append(queue, node.Right)
			}
			idx++
		}
	}
	return root
}

func serializeTree(root *TreeNode) []*int {
	if root == nil {
		return nil
	}
	var out []*int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, intPtr(node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}
	return out
}
