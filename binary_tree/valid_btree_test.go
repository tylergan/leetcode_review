package binary_tree

import "testing"

func TestIsValidBST(t *testing.T) {
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
			input: []*int{intPtr(2)},
			want:  true,
		},
		{
			name:  "example valid",
			input: []*int{intPtr(2), intPtr(1), intPtr(3)},
			want:  true,
		},
		{
			name:  "example invalid immediate",
			input: []*int{intPtr(1), intPtr(2), intPtr(3)},
			want:  false,
		},
		{
			name:  "valid right subtree",
			input: []*int{intPtr(5), intPtr(1), intPtr(7), nil, nil, intPtr(6), intPtr(8)},
			want:  true,
		},
		{
			name:  "invalid violates ancestor",
			input: []*int{intPtr(5), intPtr(1), intPtr(7), nil, nil, intPtr(4), intPtr(8)},
			want:  false,
		},
		{
			name:  "duplicates not allowed",
			input: []*int{intPtr(2), intPtr(2), intPtr(3)},
			want:  false,
		},
		{
			name:  "negative values valid",
			input: []*int{intPtr(0), intPtr(-3), intPtr(9), intPtr(-10), nil, intPtr(5)},
			want:  true,
		},
		{
			name:  "right skewed valid",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  true,
		},
		{
			name:  "left skewed valid",
			input: []*int{intPtr(4), intPtr(3), nil, intPtr(2), nil, intPtr(1)},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := isValidBST(root); got != tt.want {
				t.Fatalf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
