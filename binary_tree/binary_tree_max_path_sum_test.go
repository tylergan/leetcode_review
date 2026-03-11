package binary_tree

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  int
	}{
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  1,
		},
		{
			name:  "single negative node",
			input: []*int{intPtr(-3)},
			want:  -3,
		},
		{
			name:  "example 1 - path through root",
			input: []*int{intPtr(1), intPtr(2), intPtr(3)},
			want:  6,
		},
		{
			name:  "example 2 - path not through root",
			input: []*int{intPtr(-15), intPtr(10), intPtr(20), nil, nil, intPtr(15), intPtr(5), intPtr(-5)},
			want:  40,
		},
		{
			name:  "all negative - pick least negative",
			input: []*int{intPtr(-10), intPtr(-20), intPtr(-3)},
			want:  -3,
		},
		{
			name:  "negative children ignored",
			input: []*int{intPtr(5), intPtr(-2), intPtr(-1)},
			want:  5,
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
			want:  6,
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3)},
			want:  6,
		},
		{
			name:  "best path in subtree",
			input: []*int{intPtr(-1), intPtr(5), nil, intPtr(4), nil, intPtr(2), intPtr(1)},
			want:  11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := maxPathSum(root); got != tt.want {
				t.Fatalf("maxPathSum() = %d, want %d", got, tt.want)
			}
		})
	}
}
