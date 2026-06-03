package graphs

import "testing"

func TestValidTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{
			name:  "example 1 - connected acyclic graph",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			want:  true,
		},
		{
			name:  "example 2 - cycle",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			want:  false,
		},
		{
			name:  "single node with no edges",
			n:     1,
			edges: [][]int{},
			want:  true,
		},
		{
			name:  "disconnected graph with too few edges",
			n:     4,
			edges: [][]int{{0, 1}, {2, 3}},
			want:  false,
		},
		{
			name:  "connected graph with too many edges",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}},
			want:  false,
		},
		{
			name:  "too few edges even without cycle",
			n:     3,
			edges: [][]int{{0, 1}},
			want:  false,
		},
		{
			name:  "simple chain",
			n:     4,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validTree(tt.n, tt.edges)
			if got != tt.want {
				t.Fatalf("validTree(%d, %v) = %v, want %v", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
