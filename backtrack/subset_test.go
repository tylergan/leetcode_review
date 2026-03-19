package backtrack

import (
	"fmt"
	"slices"
	"testing"
)

func sortSubsets(ss [][]int) {
	slices.SortFunc(ss, func(a, b []int) int {
		for i := range min(len(a), len(b)) {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return len(a) - len(b)
	})
}

func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	sortSubsets(a)
	sortSubsets(b)
	for i := range a {
		if !slices.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1 - three elements",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "example 2 - single element",
			nums: []int{7},
			want: [][]int{{}, {7}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "four elements",
			nums: []int{1, 2, 3, 4},
			want: [][]int{
				{}, {1}, {2}, {3}, {4},
				{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
				{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4},
				{1, 2, 3, 4},
			},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 0, 1},
			want: [][]int{{}, {-1}, {0}, {1}, {-1, 0}, {-1, 1}, {0, 1}, {-1, 0, 1}},
		},
		{
			name: "total count is 2^n",
			nums: []int{5, 10, 15},
			want: [][]int{{}, {5}, {10}, {15}, {5, 10}, {5, 15}, {10, 15}, {5, 10, 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if !equalSubsets(got, tt.want) {
				t.Fatalf("subsets(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSubsetsCount(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1}, want: 2},
		{nums: []int{1, 2}, want: 4},
		{nums: []int{1, 2, 3}, want: 8},
		{nums: []int{1, 2, 3, 4}, want: 16},
		{nums: []int{1, 2, 3, 4, 5}, want: 32},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", len(tt.nums)), func(t *testing.T) {
			got := len(subsets(tt.nums))
			if got != tt.want {
				t.Fatalf("len(subsets(%v)) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
