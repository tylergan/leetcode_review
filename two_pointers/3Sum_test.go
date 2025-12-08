package two_pointers

import (
	"fmt"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1: mixed with duplicates",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2: no solution",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3: all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "minimum size - no solution",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "minimum size - with solution",
			nums: []int{-1, 0, 1},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "all negative",
			nums: []int{-5, -4, -3, -2, -1},
			want: [][]int{},
		},
		{
			name: "all positive",
			nums: []int{1, 2, 3, 4, 5},
			want: [][]int{},
		},
		{
			name: "multiple solutions",
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name: "many duplicates - single solution",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "duplicates with different triplets",
			nums: []int{-1, -1, -1, 0, 0, 1, 1, 2},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "large negative and positive",
			nums: []int{-100000, 0, 100000},
			want: [][]int{{-100000, 0, 100000}},
		},
		{
			name: "two pairs of duplicates",
			nums: []int{-2, -2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "complex case with multiple solutions",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
		{
			name: "single zero with negatives and positives",
			nums: []int{-5, -3, 0, 3, 5},
			want: [][]int{{-5, 0, 5}, {-3, 0, 3}},
		},
		{
			name: "multiple zeros with numbers",
			nums: []int{-1, 0, 0, 0, 1},
			want: [][]int{{-1, 0, 1}, {0, 0, 0}},
		},
		{
			name: "consecutive numbers",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			want: [][]int{{-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "no duplicates simple case",
			nums: []int{-2, -1, 0, 1, 2},
			want: [][]int{{-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "larger array",
			nums: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
			want: [][]int{{-5, 0, 5}, {-5, 1, 4}, {-5, 2, 3}, {-4, -1, 5}, {-4, 0, 4}, {-4, 1, 3}, {-3, -2, 5}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "same number three times",
			nums: []int{-3, -3, 1, 2, 3, 3, 6},
			want: [][]int{{-3, -3, 6}, {-3, 1, 2}},
		},
		{
			name: "edge case - barely sums to zero",
			nums: []int{-1, -1, 2},
			want: [][]int{{-1, -1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !equalTriplets(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equalTriplets compares two slices of triplets for equality,
// ignoring the order of triplets and the order within each triplet
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert triplets to sorted string representations for comparison
	aSet := make(map[string]bool)
	bSet := make(map[string]bool)

	for _, triplet := range a {
		sorted := make([]int, len(triplet))
		copy(sorted, triplet)
		sort.Ints(sorted)
		key := tripletToString(sorted)
		aSet[key] = true
	}

	for _, triplet := range b {
		sorted := make([]int, len(triplet))
		copy(sorted, triplet)
		sort.Ints(sorted)
		key := tripletToString(sorted)
		bSet[key] = true
	}

	if len(aSet) != len(bSet) {
		return false
	}

	for key := range aSet {
		if !bSet[key] {
			return false
		}
	}

	return true
}

// tripletToString converts a sorted triplet to a string for comparison
func tripletToString(triplet []int) string {
	return fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
}

func TestThreeSumProperties(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(t *testing.T, result [][]int)
	}{
		{
			name: "no duplicate triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			validate: func(t *testing.T, result [][]int) {
				seen := make(map[string]bool)
				for _, triplet := range result {
					sorted := make([]int, len(triplet))
					copy(sorted, triplet)
					sort.Ints(sorted)
					key := tripletToString(sorted)
					if seen[key] {
						t.Errorf("Duplicate triplet found: %v", triplet)
					}
					seen[key] = true
				}
			},
		},
		{
			name: "all triplets sum to zero",
			nums: []int{-2, -1, 0, 1, 2, 3},
			validate: func(t *testing.T, result [][]int) {
				for _, triplet := range result {
					if len(triplet) != 3 {
						t.Errorf("Triplet has wrong size: %v", triplet)
						continue
					}
					sum := triplet[0] + triplet[1] + triplet[2]
					if sum != 0 {
						t.Errorf("Triplet %v does not sum to 0, sum = %d", triplet, sum)
					}
				}
			},
		},
		{
			name: "triplets are sorted",
			nums: []int{-4, -1, 0, 1, 2, 3},
			validate: func(t *testing.T, result [][]int) {
				for _, triplet := range result {
					if len(triplet) != 3 {
						continue
					}
					if triplet[0] > triplet[1] || triplet[1] > triplet[2] {
						t.Errorf("Triplet is not sorted: %v", triplet)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			tt.validate(t, result)
		})
	}
}
