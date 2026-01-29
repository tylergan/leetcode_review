package binary_tree

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  [][]int
	}{
		{
			name:  "empty",
			input: nil,
			want:  [][]int{},
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  [][]int{{1}},
		},
		{
			name:  "perfect tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
		{
			name:  "example with gap",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), nil, intPtr(7)},
			want:  [][]int{{1}, {2, 3}, {4, 5, 7}},
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name:  "negative values",
			input: []*int{intPtr(0), intPtr(-1), intPtr(2), intPtr(-3), nil, nil, intPtr(4)},
			want:  [][]int{{0}, {-1, 2}, {-3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			got := levelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
