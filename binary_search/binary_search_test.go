package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example 1: target found",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 4,
			want:   3,
		},
		{
			name:   "example 2: target missing",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 3,
			want:   -1,
		},
		{
			name:   "target at first element",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			name:   "target at last element",
			nums:   []int{1, 3},
			target: 3,
			want:   1,
		},
		{
			name:   "single element found",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element missing",
			nums:   []int{5},
			target: 4,
			want:   -1,
		},
		{
			name:   "negative numbers",
			nums:   []int{-5, -3, -1},
			target: -3,
			want:   1,
		},
		{
			name:   "target below minimum",
			nums:   []int{2, 4, 6},
			target: 1,
			want:   -1,
		},
		{
			name:   "target above maximum",
			nums:   []int{2, 4, 6},
			target: 7,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
