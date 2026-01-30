package binary_tree

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  []int
	}{
		{
			name:  "empty",
			input: nil,
			want:  []int{},
		},
		{
			name:  "single node",
			input: []*int{intPtr(1)},
			want:  []int{1},
		},
		{
			name:  "example small",
			input: []*int{intPtr(1), intPtr(2), intPtr(3)},
			want:  []int{1, 3},
		},
		{
			name:  "perfect tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  []int{1, 3, 7},
		},
		{
			name:  "mixed gaps",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(5), nil, intPtr(4)},
			want:  []int{1, 3, 4},
		},
		{
			name:  "left skewed",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "right skewed",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "left nodes hidden by right",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), nil, nil, intPtr(5)},
			want:  []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			got := rightSideView(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
