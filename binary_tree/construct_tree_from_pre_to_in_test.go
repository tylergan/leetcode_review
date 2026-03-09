package binary_tree

import (
	"reflect"
	"testing"
)

func TestBuildTreeFromPreorderInorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []*int
	}{
		{
			name:     "single node",
			preorder: []int{1},
			inorder:  []int{1},
			want:     []*int{intPtr(1)},
		},
		{
			name:     "leetcode example",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)},
		},
		{
			name:     "problem example",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{2, 1, 3, 4},
			want:     []*int{intPtr(1), intPtr(2), intPtr(3), nil, nil, nil, intPtr(4)},
		},
		{
			name:     "left skewed tree",
			preorder: []int{3, 2, 1},
			inorder:  []int{1, 2, 3},
			want:     []*int{intPtr(3), intPtr(2), nil, intPtr(1)},
		},
		{
			name:     "right skewed tree",
			preorder: []int{1, 2, 3},
			inorder:  []int{1, 2, 3},
			want:     []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3)},
		},
		{
			name:     "sparse asymmetric tree",
			preorder: []int{1, 2, 4, 3, 5},
			inorder:  []int{2, 4, 1, 5, 3},
			want:     []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(4), intPtr(5)},
		},
		{
			name:     "negative values",
			preorder: []int{0, -3, -10, 9, 5},
			inorder:  []int{-10, -3, 0, 5, 9},
			want:     []*int{intPtr(0), intPtr(-3), intPtr(9), intPtr(-10), nil, intPtr(5)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreeFromPreIn(tt.preorder, tt.inorder)
			gotLevel := serializeTree(got)
			if !reflect.DeepEqual(gotLevel, tt.want) {
				t.Fatalf("buildTreeFromPreIn() = %v, want %v", gotLevel, tt.want)
			}
		})
	}
}
