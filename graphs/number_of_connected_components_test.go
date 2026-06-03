package graphs

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "example 1 - two components",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			want:  2,
		},
		{
			name:  "example 2 - one component",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:  1,
		},
		{
			name:  "all isolated nodes",
			n:     4,
			edges: [][]int{},
			want:  4,
		},
		{
			name:  "cycle still one component",
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
			want:  1,
		},
		{
			name:  "component plus isolated node",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}},
			want:  2,
		},
		{
			name:  "multiple separate pairs",
			n:     6,
			edges: [][]int{{0, 1}, {2, 3}, {4, 5}},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countComponents(tt.n, tt.edges)
			if got != tt.want {
				t.Fatalf("countComponents(%d, %v) = %d, want %d", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
