package arrays

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "example 1: basic case",
			nums:   []int{3, 4, 5, 6},
			target: 7,
			want:   []int{0, 1},
		},
		{
			name:   "example 2: non-adjacent indices",
			nums:   []int{4, 5, 6},
			target: 10,
			want:   []int{0, 2},
		},
		{
			name:   "example 3: duplicate values",
			nums:   []int{5, 5},
			target: 10,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "mix of positive and negative",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "with zeros",
			nums:   []int{0, 4, 3, 0},
			target: 0,
			want:   []int{0, 3},
		},
		{
			name:   "large numbers",
			nums:   []int{1000000, 2000000, 3000000},
			target: 5000000,
			want:   []int{1, 2},
		},
		{
			name:   "target at end of array",
			nums:   []int{1, 2, 3, 4, 5},
			target: 9,
			want:   []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != 2 {
				t.Errorf("twoSum() returned slice of length %d, want 2", len(got))
				return
			}
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
