package stack

import "testing"

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			name:     "example 1",
			target:   10,
			position: []int{1, 4},
			speed:    []int{3, 2},
			want:     1,
		},
		{
			name:     "example 2",
			target:   10,
			position: []int{4, 1, 0, 7},
			speed:    []int{2, 2, 1, 1},
			want:     3,
		},
		{
			name:     "neetcode classic",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			name:     "single car",
			target:   100,
			position: []int{0},
			speed:    []int{1},
			want:     1,
		},
		{
			name:     "all separate fleets",
			target:   10,
			position: []int{6, 8},
			speed:    []int{3, 2},
			want:     2,
		},
		{
			name:     "all merge into one",
			target:   10,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
		{
			name:     "meet at destination",
			target:   10,
			position: []int{0, 4},
			speed:    []int{2, 1},
			want:     1,
		},
		{
			name:     "fast behind slow ahead",
			target:   15,
			position: []int{5, 10},
			speed:    []int{4, 1},
			want:     1,
		},
		{
			name:     "already ordered positions",
			target:   20,
			position: []int{2, 4, 6, 8},
			speed:    []int{4, 3, 2, 1},
			want:     1,
		},
		{
			name:     "no merges due to equal times",
			target:   10,
			position: []int{0, 2, 4},
			speed:    []int{1, 1, 1},
			want:     3,
		},
		{
			name:   "all merged bounded by the final car fleet",
			target: 10,
			// initially car at position 2 is faster than position 1, but once it bounds
			// into final car fleet, the first one eventually catches up
			position: []int{0, 4, 2},
			speed:    []int{2, 1, 3},
			want:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := carFleet(tt.target, tt.position, tt.speed)
			if got != tt.want {
				t.Errorf("carFleet(%d, %v, %v) = %d, want %d", tt.target, tt.position, tt.speed, got, tt.want)
			}
		})
	}
}
