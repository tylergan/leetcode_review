package binary_tree

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  bool
	}{
		{
			name:  "empty",
			input: nil,
			want:  true,
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  true,
		},
		{
			name:  "balanced example",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, nil, intPtr(4)},
			want:  true,
		},
		{
			name:  "unbalanced example",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, nil, intPtr(4), nil, intPtr(5)},
			want:  false,
		},
		{
			name:  "perfectly balanced",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  true,
		},
		{
			name:  "left heavy chain",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  false,
		},
		{
			name:  "right heavy chain",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  false,
		},
		{
			name:  "balanced with missing leaves",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), nil, nil, intPtr(5)},
			want:  true,
		},
		{
			name:  "local imbalance",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), nil, nil, intPtr(6)},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := isBalanced(root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
