package heappq

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 1, 5, 4},
			k:    2,
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{2, 3, 1, 1, 5, 5, 4},
			k:    3,
			want: 4,
		},
		{
			name: "k equals 1 - max element",
			nums: []int{3, 1, 2},
			k:    1,
			want: 3,
		},
		{
			name: "k equals length - min element",
			nums: []int{3, 1, 2},
			k:    3,
			want: 1,
		},
		{
			name: "single element",
			nums: []int{7},
			k:    1,
			want: 7,
		},
		{
			name: "all duplicates",
			nums: []int{5, 5, 5, 5},
			k:    2,
			want: 5,
		},
		{
			name: "negative values",
			nums: []int{-1, -5, -3, -2, -4},
			k:    2,
			want: -2,
		},
		{
			name: "mixed positive and negative",
			nums: []int{-1, 2, 0, -3, 4},
			k:    3,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.nums, tt.k); got != tt.want {
				t.Fatalf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
