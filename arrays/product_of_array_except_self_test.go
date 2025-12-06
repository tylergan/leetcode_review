package arrays

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1: basic positive numbers",
			nums: []int{1, 2, 4, 6},
			want: []int{48, 24, 12, 8},
		},
		{
			name: "example 2: with single zero",
			nums: []int{-1, 0, 1, 2, 3},
			want: []int{0, -6, 0, 0, 0},
		},
		{
			name: "minimum size array",
			nums: []int{2, 3},
			want: []int{3, 2},
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "with negative numbers",
			nums: []int{-1, 2, -3, 4},
			want: []int{-24, 12, -8, 6},
		},
		{
			name: "all negative numbers",
			nums: []int{-2, -3, -4},
			want: []int{12, 8, 6},
		},
		{
			name: "zero at beginning",
			nums: []int{0, 2, 3, 4},
			want: []int{24, 0, 0, 0},
		},
		{
			name: "zero at end",
			nums: []int{2, 3, 4, 0},
			want: []int{0, 0, 0, 24},
		},
		{
			name: "zero in middle",
			nums: []int{2, 0, 4, 5},
			want: []int{0, 40, 0, 0},
		},
		{
			name: "multiple zeros",
			nums: []int{0, 0, 1, 2},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "contains one",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "large products",
			nums: []int{10, 10, 10},
			want: []int{100, 100, 100},
		},
		{
			name: "mix with positive one and negative",
			nums: []int{1, -1, 2},
			want: []int{-2, 2, -1},
		},
		{
			name: "negative zero boundary",
			nums: []int{-20, 0, 20},
			want: []int{0, -400, 0},
		},
		{
			name: "single negative with positives",
			nums: []int{2, 3, -4, 5},
			want: []int{-60, -40, 30, -24},
		},
		{
			name: "two elements with zero",
			nums: []int{0, 5},
			want: []int{5, 0},
		},
		{
			name: "two elements both zero",
			nums: []int{0, 0},
			want: []int{0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductExceptSelfProperties(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		validate func(t *testing.T, nums, result []int)
	}{
		{
			name: "result length matches input length",
			nums: []int{1, 2, 3, 4, 5},
			validate: func(t *testing.T, nums, result []int) {
				if len(result) != len(nums) {
					t.Errorf("Result length %d does not match input length %d", len(result), len(nums))
				}
			},
		},
		{
			name: "no zeros - all products are non-zero",
			nums: []int{1, 2, 3, 4},
			validate: func(t *testing.T, nums, result []int) {
				for i, val := range result {
					if val == 0 {
						t.Errorf("Expected non-zero result at index %d, got 0", i)
					}
				}
			},
		},
		{
			name: "single zero - exactly one non-zero result",
			nums: []int{1, 0, 2, 3},
			validate: func(t *testing.T, nums, result []int) {
				nonZeroCount := 0
				for _, val := range result {
					if val != 0 {
						nonZeroCount++
					}
				}
				if nonZeroCount != 1 {
					t.Errorf("Expected exactly 1 non-zero result with single zero, got %d", nonZeroCount)
				}
			},
		},
		{
			name: "multiple zeros - all results are zero",
			nums: []int{1, 0, 2, 0, 3},
			validate: func(t *testing.T, nums, result []int) {
				for i, val := range result {
					if val != 0 {
						t.Errorf("Expected zero result at index %d with multiple zeros, got %d", i, val)
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			tt.validate(t, tt.nums, result)
		})
	}
}
