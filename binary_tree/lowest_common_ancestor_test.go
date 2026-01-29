package binary_tree

import "testing"

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		pVal  int
		qVal  int
		want  int
	}{
		{
			name:  "example root",
			input: []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(4), intPtr(7), intPtr(9), nil, intPtr(2)},
			pVal:  3,
			qVal:  8,
			want:  5,
		},
		{
			name:  "example ancestor",
			input: []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(4), intPtr(7), intPtr(9), nil, intPtr(2)},
			pVal:  3,
			qVal:  4,
			want:  3,
		},
		{
			name:  "same subtree",
			input: []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(4), intPtr(7), intPtr(9), nil, intPtr(2)},
			pVal:  7,
			qVal:  9,
			want:  8,
		},
		{
			name:  "descendant is lca",
			input: []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(1), intPtr(4), intPtr(7), intPtr(9), nil, intPtr(2)},
			pVal:  1,
			qVal:  2,
			want:  1,
		},
		{
			name:  "negative values",
			input: []*int{intPtr(0), intPtr(-3), intPtr(9), intPtr(-10), nil, intPtr(5)},
			pVal:  -10,
			qVal:  5,
			want:  0,
		},
		{
			name:  "unordered inputs",
			input: []*int{intPtr(0), intPtr(-3), intPtr(9), intPtr(-10), nil, intPtr(5)},
			pVal:  9,
			qVal:  -3,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			if p == nil || q == nil {
				t.Fatalf("failed to locate nodes p=%d q=%d", tt.pVal, tt.qVal)
			}
			got := lowestCommonAncestor(root, p, q)
			if got == nil {
				t.Fatalf("lowestCommonAncestor() returned nil")
			}
			if got.Val != tt.want {
				t.Fatalf("lowestCommonAncestor() = %d, want %d", got.Val, tt.want)
			}
		})
	}
}
