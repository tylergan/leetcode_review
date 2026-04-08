package backtrack

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1 - duplicates present",
			nums: []int{1, 2, 2},
			want: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
		{
			name: "example 2 - no duplicates",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
		{
			name: "all duplicates",
			nums: []int{3, 3, 3},
			want: [][]int{{}, {3}, {3, 3}, {3, 3, 3}},
		},
		{
			name: "no duplicates behaves like subsets I",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "two pairs of duplicates",
			nums: []int{1, 2, 1, 2},
			want: [][]int{
				{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2},
				{1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2},
			},
		},
		{
			name: "single element",
			nums: []int{5},
			want: [][]int{{}, {5}},
		},
		{
			name: "negative numbers with duplicates",
			nums: []int{-1, -1, 2},
			want: [][]int{{}, {-1}, {-1, -1}, {-1, 2}, {-1, -1, 2}, {2}},
		},
		{
			name: "already sorted input",
			nums: []int{1, 1, 2, 2},
			want: [][]int{
				{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2},
				{1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2},
			},
		},
		{
			name: "unsorted input produces same result",
			nums: []int{2, 1, 2, 1},
			want: [][]int{
				{}, {1}, {2}, {1, 1}, {1, 2}, {2, 2},
				{1, 1, 2}, {1, 2, 2}, {1, 1, 2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			if !equalSubsets(got, tt.want) {
				t.Fatalf("subsetsWithDup(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestSubsetsWithDupNoDuplicateSubsets(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{nums: []int{1, 2, 2}},
		{nums: []int{1, 1, 1, 1}},
		{nums: []int{1, 2, 1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%v", tt.nums), func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			seen := map[string]bool{}
			for _, subset := range got {
				key := fmt.Sprintf("%v", subset)
				if seen[key] {
					t.Fatalf("duplicate subset %v in result %v", subset, got)
				}
				seen[key] = true
			}
		})
	}
}

func TestSubsetsWithDupCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "no dups n=3", nums: []int{1, 2, 3}, want: 8},
		{name: "all same n=3", nums: []int{2, 2, 2}, want: 4},
		{name: "one dup", nums: []int{1, 2, 2}, want: 6},
		{name: "single", nums: []int{1}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(subsetsWithDup(tt.nums))
			if got != tt.want {
				t.Fatalf("len(subsetsWithDup(%v)) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
