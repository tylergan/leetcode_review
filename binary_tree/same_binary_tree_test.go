package binary_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []*int
		q    []*int
		want bool
	}{
		{
			name: "both empty",
			p:    nil,
			q:    nil,
			want: true,
		},
		{
			name: "one empty",
			p:    []*int{intPtr(1)},
			q:    nil,
			want: false,
		},
		{
			name: "same structure values",
			p:    []*int{intPtr(1), intPtr(2), intPtr(3)},
			q:    []*int{intPtr(1), intPtr(2), intPtr(3)},
			want: true,
		},
		{
			name: "different structure",
			p:    []*int{intPtr(4), intPtr(7)},
			q:    []*int{intPtr(4), nil, intPtr(7)},
			want: false,
		},
		{
			name: "different values",
			p:    []*int{intPtr(1), intPtr(2), intPtr(3)},
			q:    []*int{intPtr(1), intPtr(3), intPtr(2)},
			want: false,
		},
		{
			name: "deep mismatch",
			p:    []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(4)},
			q:    []*int{intPtr(1), intPtr(2), intPtr(3), nil, intPtr(5)},
			want: false,
		},
		{
			name: "both skewed same",
			p:    []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
			q:    []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := buildTree(tt.p)
			q := buildTree(tt.q)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
