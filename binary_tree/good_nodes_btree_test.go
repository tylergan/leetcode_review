package binary_tree

import "testing"

func TestGoodNodes(t *testing.T) {
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
			name:  "example 1",
			input: []*int{intPtr(2), intPtr(1), intPtr(1), intPtr(3), nil, intPtr(1), intPtr(5)},
			want:  3,
		},
		{
			name:  "example 2",
			input: []*int{intPtr(1), intPtr(2), intPtr(-1), intPtr(3), intPtr(4)},
			want:  4,
		},
		{
			name:  "all increasing",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  4,
		},
		{
			name:  "all decreasing",
			input: []*int{intPtr(4), intPtr(3), nil, intPtr(2), nil, intPtr(1)},
			want:  1,
		},
		{
			name:  "mixed with negatives",
			input: []*int{intPtr(-2), intPtr(-3), intPtr(-1), intPtr(-4), nil, nil, intPtr(0)},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := goodNodes(root); got != tt.want {
				t.Fatalf("goodNodes() = %d, want %d", got, tt.want)
			}
		})
	}
}
