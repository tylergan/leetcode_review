package binary_tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  int
	}{
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  0,
		},
		{
			name:  "two nodes",
			input: []*int{intPtr(1), intPtr(2)},
			want:  1,
		},
		{
			name:  "example",
			input: []*int{intPtr(1), nil, intPtr(2), intPtr(3), intPtr(4), intPtr(5)},
			want:  3,
		},
		{
			name:  "balanced",
			input: []*int{intPtr(1), intPtr(2), intPtr(3)},
			want:  2,
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  3,
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  3,
		},
		{
			name:  "diameter not through root",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), intPtr(4), nil, nil, intPtr(5)},
			want:  3,
		},
		{
			name:  "single chain with extra leaf",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4), nil, intPtr(5), intPtr(6)},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := diameterOfBinaryTree(root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %d, want %d", got, tt.want)
			}
		})
	}
}
