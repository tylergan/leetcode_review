package graphs

import (
	"slices"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{
			name:  "example 1 - last edge closes cycle",
			edges: [][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}},
			want:  []int{2, 4},
		},
		{
			name:  "example 2 - inner cycle",
			edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}},
			want:  []int{3, 4},
		},
		{
			name:  "triangle",
			edges: [][]int{{1, 2}, {2, 3}, {1, 3}},
			want:  []int{1, 3},
		},
		{
			name:  "returns last redundant edge encountered",
			edges: [][]int{{1, 2}, {2, 3}, {3, 1}, {4, 5}, {5, 1}, {2, 4}},
			want:  []int{2, 4},
		},
		{
			name:  "long chain closed at end",
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}},
			want:  []int{1, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantConnection(tt.edges)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("findRedundantConnection(%v) = %v, want %v", tt.edges, got, tt.want)
			}
		})
	}
}
