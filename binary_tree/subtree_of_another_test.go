package binary_tree

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    []*int
		subRoot []*int
		want    bool
	}{
		{
			name:    "subtree is root",
			root:    []*int{intPtr(1), intPtr(2), intPtr(3)},
			subRoot: []*int{intPtr(1), intPtr(2), intPtr(3)},
			want:    true,
		},
		{
			name:    "example true",
			root:    []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5)},
			subRoot: []*int{intPtr(2), intPtr(4), intPtr(5)},
			want:    true,
		},
		{
			name:    "example false",
			root:    []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), nil, nil, intPtr(6)},
			subRoot: []*int{intPtr(2), intPtr(4), intPtr(5)},
			want:    false,
		},
		{
			name:    "subtree at leaf",
			root:    []*int{intPtr(5), intPtr(3), intPtr(8), intPtr(2), intPtr(4), intPtr(7), intPtr(9)},
			subRoot: []*int{intPtr(2)},
			want:    true,
		},
		{
			name:    "duplicate values",
			root:    []*int{intPtr(1), intPtr(1), intPtr(1), intPtr(1), nil, nil, intPtr(1)},
			subRoot: []*int{intPtr(1), nil, intPtr(1)},
			want:    true,
		},
		{
			name:    "right skewed match",
			root:    []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			subRoot: []*int{intPtr(2), nil, intPtr(3)},
			want:    false,
		},
		{
			name:    "missing subtree",
			root:    []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), nil, nil, intPtr(5)},
			subRoot: []*int{intPtr(2), nil, intPtr(4)},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			subRoot := buildTree(tt.subRoot)
			if got := isSubtree(root, subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
