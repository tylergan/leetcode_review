package two_pointers

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example 1: basic case",
			numbers: []int{1, 2, 3, 4},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "minimum array size",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "two same numbers",
			numbers: []int{1, 1},
			target:  2,
			want:    []int{1, 2},
		},
		{
			name:    "sum at beginning of array",
			numbers: []int{1, 2, 5, 10, 20},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "sum at end of array",
			numbers: []int{1, 2, 5, 10, 20},
			target:  30,
			want:    []int{4, 5},
		},
		{
			name:    "sum in middle",
			numbers: []int{1, 2, 5, 10, 20},
			target:  15,
			want:    []int{3, 4},
		},
		{
			name:    "first and last elements",
			numbers: []int{1, 5, 10, 15, 20},
			target:  21,
			want:    []int{1, 5},
		},
		{
			name:    "negative numbers",
			numbers: []int{-10, -5, -4, 0, 3, 5},
			target:  -7,
			want:    []int{1, 5},
		},
		{
			name:    "all negative numbers",
			numbers: []int{-10, -8, -5, -3, -1},
			target:  -13,
			want:    []int{1, 4},
		},
		{
			name:    "negative and positive sum to zero",
			numbers: []int{-7, -3, 0, 2, 7},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "target is zero with negatives",
			numbers: []int{-10, -3, 0, 5, 10},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "contains zeros",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "zero plus positive",
			numbers: []int{0, 1, 2, 4},
			target:  4,
			want:    []int{1, 4},
		},
		{
			name:    "adjacent elements",
			numbers: []int{1, 3, 4, 5, 7},
			target:  7,
			want:    []int{2, 3},
		},
		{
			name:    "large span",
			numbers: []int{1, 2, 3, 10, 11, 12},
			target:  13,
			want:    []int{1, 6},
		},
		{
			name:    "duplicates in array",
			numbers: []int{1, 2, 2, 3, 5},
			target:  7,
			want:    []int{2, 5},
		},
		{
			name:    "negative target",
			numbers: []int{-100, -50, -10, -5},
			target:  -110,
			want:    []int{1, 3},
		},
		{
			name:    "large positive numbers",
			numbers: []int{100, 200, 300, 400, 500},
			target:  900,
			want:    []int{4, 5},
		},
		{
			name:    "maximum constraints",
			numbers: []int{-1000, -500, 0, 600, 1000},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "sequential numbers",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:  19,
			want:    []int{9, 10},
		},
		{
			name:    "two negative sum to positive",
			numbers: []int{-10, -5, -3, 8, 15},
			target:  -8,
			want:    []int{2, 3},
		},
		{
			name:    "mix with zero target",
			numbers: []int{-5, -3, 0, 2, 5},
			target:  0,
			want:    []int{1, 5},
		},
		{
			name:    "elements far apart",
			numbers: []int{1, 100, 200, 300, 399},
			target:  400,
			want:    []int{1, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwoSumProperties(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		validate func(t *testing.T, numbers []int, target int, result []int)
	}{
		{
			name:    "indices are 1-indexed",
			numbers: []int{1, 2, 3, 4},
			target:  3,
			validate: func(t *testing.T, numbers []int, target int, result []int) {
				if len(result) != 2 {
					t.Errorf("Expected 2 indices, got %d", len(result))
					return
				}
				if result[0] < 1 || result[1] < 1 {
					t.Errorf("Indices should be 1-indexed, got %v", result)
				}
			},
		},
		{
			name:    "index1 < index2",
			numbers: []int{1, 2, 3, 4},
			target:  5,
			validate: func(t *testing.T, numbers []int, target int, result []int) {
				if len(result) != 2 {
					t.Errorf("Expected 2 indices, got %d", len(result))
					return
				}
				if result[0] >= result[1] {
					t.Errorf("index1 should be < index2, got %v", result)
				}
			},
		},
		{
			name:    "sum equals target",
			numbers: []int{1, 2, 3, 4},
			target:  7,
			validate: func(t *testing.T, numbers []int, target int, result []int) {
				if len(result) != 2 {
					t.Errorf("Expected 2 indices, got %d", len(result))
					return
				}
				// Convert to 0-indexed for array access
				idx1, idx2 := result[0]-1, result[1]-1
				sum := numbers[idx1] + numbers[idx2]
				if sum != target {
					t.Errorf("Sum of elements at indices %v = %d, want %d", result, sum, target)
				}
			},
		},
		{
			name:    "indices within bounds",
			numbers: []int{1, 2, 3, 4, 5},
			target:  9,
			validate: func(t *testing.T, numbers []int, target int, result []int) {
				if len(result) != 2 {
					t.Errorf("Expected 2 indices, got %d", len(result))
					return
				}
				if result[0] < 1 || result[0] > len(numbers) {
					t.Errorf("index1 out of bounds: %d", result[0])
				}
				if result[1] < 1 || result[1] > len(numbers) {
					t.Errorf("index2 out of bounds: %d", result[1])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.numbers, tt.target)
			tt.validate(t, tt.numbers, tt.target, result)
		})
	}
}
