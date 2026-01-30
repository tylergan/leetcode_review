package binary_tree

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		k     int
		want  int
	}{
		{
			name:  "example 1",
			input: []*int{intPtr(2), intPtr(1), intPtr(3)},
			k:     1,
			want:  1,
		},
		{
			name:  "example 2",
			input: []*int{intPtr(4), intPtr(3), intPtr(5), intPtr(2), nil},
			k:     4,
			want:  5,
		},
		{
			name:  "single node",
			input: []*int{intPtr(7)},
			k:     1,
			want:  7,
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(5), intPtr(4), nil, intPtr(3), nil, intPtr(2), nil, intPtr(1)},
			k:     3,
			want:  3,
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4), nil, intPtr(5)},
			k:     5,
			want:  5,
		},
		{
			name:  "balanced tree middle",
			input: []*int{intPtr(5), intPtr(3), intPtr(7), intPtr(2), intPtr(4), intPtr(6), intPtr(8)},
			k:     4,
			want:  5,
		},
		{
			name:  "balanced tree smallest",
			input: []*int{intPtr(5), intPtr(3), intPtr(7), intPtr(2), intPtr(4), intPtr(6), intPtr(8)},
			k:     1,
			want:  2,
		},
		{
			name:  "balanced tree largest",
			input: []*int{intPtr(5), intPtr(3), intPtr(7), intPtr(2), intPtr(4), intPtr(6), intPtr(8)},
			k:     7,
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			if got := kthSmallest(root, tt.k); got != tt.want {
				t.Fatalf("kthSmallest() = %d, want %d", got, tt.want)
			}
		})
	}
}
