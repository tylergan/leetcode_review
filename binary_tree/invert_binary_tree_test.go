package binary_tree

import (
	"reflect"
	"testing"
)

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

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  []*int
	}{
		{
			name:  "empty",
			input: nil,
			want:  nil,
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  []*int{intPtr(1)},
		},
		{
			name:  "perfect tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  []*int{intPtr(1), intPtr(3), intPtr(2), intPtr(7), intPtr(6), intPtr(5), intPtr(4)},
		},
		{
			name:  "three nodes",
			input: []*int{intPtr(3), intPtr(2), intPtr(1)},
			want:  []*int{intPtr(3), intPtr(1), intPtr(2)},
		},
		{
			name:  "sparse tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(4), nil, intPtr(5)},
			want:  []*int{intPtr(1), intPtr(3), intPtr(2), intPtr(5), nil, intPtr(4)},
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
			want:  []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			got := invertTree(root)
			gotLevel := serializeTree(got)
			if !reflect.DeepEqual(gotLevel, tt.want) {
				t.Errorf("invertTree() = %v, want %v", gotLevel, tt.want)
			}
		})
	}
}
