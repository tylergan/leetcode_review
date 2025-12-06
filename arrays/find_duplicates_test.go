package arrays

import "testing"

func TestHasDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "array with duplicates in middle",
			nums:     []int{1, 2, 3, 3},
			expected: true,
		},
		{
			name:     "array without duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element array",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "all elements are the same",
			nums:     []int{5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "duplicate at beginning",
			nums:     []int{1, 1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "duplicate at end",
			nums:     []int{1, 2, 3, 4, 4},
			expected: true,
		},
		{
			name:     "two elements same",
			nums:     []int{7, 7},
			expected: true,
		},
		{
			name:     "two elements different",
			nums:     []int{7, 8},
			expected: false,
		},
		{
			name:     "negative numbers with duplicates",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "negative numbers without duplicates",
			nums:     []int{-1, -2, -3, -4},
			expected: false,
		},
		{
			name:     "mix of positive and negative with duplicates",
			nums:     []int{-1, 0, 1, 2, -1},
			expected: true,
		},
		{
			name:     "mix of positive and negative without duplicates",
			nums:     []int{-1, 0, 1, 2, 3},
			expected: false,
		},
		{
			name:     "large array with duplicate at end",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
			expected: true,
		},
		{
			name:     "large array without duplicates",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: false,
		},
		{
			name:     "zero appears twice",
			nums:     []int{0, 1, 2, 0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("hasDuplicate(%v) = %v; expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
