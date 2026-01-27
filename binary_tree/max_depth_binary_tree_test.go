package binary_tree

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  int
	}{
		{
			name:  "empty",
			input: nil,
			want:  0,
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  1,
		},
		{
			name:  "example",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, nil, intPtr(4)},
			want:  3,
		},
		{
			name:  "perfect tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  3,
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  4,
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3)},
			want:  3,
		},
		{
			name:  "uneven depths",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), nil, nil, nil, intPtr(5)},
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %d, want %d", got, tt.want)
			}
		})
	}
}
