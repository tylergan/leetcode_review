package heappq

import (
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "example 1",
			points: [][]int{{0, 2}, {2, 2}},
			k:      1,
			want:   [][]int{{0, 2}},
		},
		{
			name:   "example 2",
			points: [][]int{{0, 2}, {2, 0}, {2, 2}},
			k:      2,
			want:   [][]int{{0, 2}, {2, 0}},
		},
		{
			name:   "single point",
			points: [][]int{{1, 1}},
			k:      1,
			want:   [][]int{{1, 1}},
		},
		{
			name:   "k equals length",
			points: [][]int{{3, 3}, {1, 1}, {2, 2}},
			k:      3,
			want:   [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
		{
			name:   "negative coordinates",
			points: [][]int{{-2, -2}, {1, 1}, {-1, 0}},
			k:      2,
			want:   [][]int{{1, 1}, {-1, 0}},
		},
		{
			name:   "origin included",
			points: [][]int{{0, 0}, {1, 1}, {2, 2}},
			k:      1,
			want:   [][]int{{0, 0}},
		},
		{
			name:   "same distance",
			points: [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			k:      2,
			want:   [][]int{{0, -1}, {0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.points, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("kClosest() returned %d points, want %d", len(got), len(tt.want))
			}
			// Sort both by (x, y) for order-independent comparison
			sortPoints(got)
			sortPoints(tt.want)
			for i := range got {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Fatalf("kClosest() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func sortPoints(pts [][]int) {
	sort.Slice(pts, func(i, j int) bool {
		if pts[i][0] != pts[j][0] {
			return pts[i][0] < pts[j][0]
		}
		return pts[i][1] < pts[j][1]
	})
}
