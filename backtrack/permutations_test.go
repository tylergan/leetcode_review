package backtrack

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1 - three elements",
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "example 2 - single element",
			nums: []int{7},
			want: [][]int{{7}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{{1, 2}, {2, 1}},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1},
			want: [][]int{
				{-1, 0, 1}, {-1, 1, 0},
				{0, -1, 1}, {0, 1, -1},
				{1, -1, 0}, {1, 0, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.nums)
			if !equalSubsets(got, tt.want) {
				t.Fatalf("permute(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestPermuteCount(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1}, want: 1},
		{nums: []int{1, 2}, want: 2},
		{nums: []int{1, 2, 3}, want: 6},
		{nums: []int{1, 2, 3, 4}, want: 24},
		{nums: []int{1, 2, 3, 4, 5}, want: 120},
		{nums: []int{1, 2, 3, 4, 5, 6}, want: 720},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", len(tt.nums)), func(t *testing.T) {
			got := len(permute(tt.nums))
			if got != tt.want {
				t.Fatalf("len(permute(%v)) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
