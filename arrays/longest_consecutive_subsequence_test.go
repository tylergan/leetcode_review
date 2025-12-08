package arrays

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1: sequence in middle",
			nums: []int{2, 20, 4, 10, 3, 4, 5},
			want: 4,
		},
		{
			name: "example 2: full sequence with duplicates",
			nums: []int{0, 3, 2, 5, 4, 6, 1, 1},
			want: 7,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "two consecutive elements",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "two non-consecutive elements",
			nums: []int{1, 3},
			want: 1,
		},
		{
			name: "all consecutive ascending",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "all consecutive descending order",
			nums: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "all consecutive random order",
			nums: []int{3, 1, 4, 2, 5},
			want: 5,
		},
		{
			name: "no consecutive - all separated",
			nums: []int{1, 3, 5, 7, 9},
			want: 1,
		},
		{
			name: "multiple duplicates",
			nums: []int{1, 2, 2, 3, 3, 3, 4},
			want: 4,
		},
		{
			name: "all same elements",
			nums: []int{5, 5, 5, 5},
			want: 1,
		},
		{
			name: "negative numbers only",
			nums: []int{-3, -1, -2, -4},
			want: 4,
		},
		{
			name: "negative and positive",
			nums: []int{-2, -1, 0, 1, 2},
			want: 5,
		},
		{
			name: "sequence crosses zero",
			nums: []int{-1, 0, 1},
			want: 3,
		},
		{
			name: "multiple sequences - different lengths",
			nums: []int{1, 2, 3, 10, 11, 12, 13},
			want: 4,
		},
		{
			name: "multiple sequences - same length",
			nums: []int{1, 2, 3, 10, 11, 12},
			want: 3,
		},
		{
			name: "sequence at start",
			nums: []int{1, 2, 3, 100, 200, 300},
			want: 3,
		},
		{
			name: "sequence at end",
			nums: []int{100, 200, 300, 1, 2, 3},
			want: 3,
		},
		{
			name: "large gaps",
			nums: []int{1, 1000000, 2, 3},
			want: 3,
		},
		{
			name: "very large numbers",
			nums: []int{1000000000, 1000000001, 1000000002},
			want: 3,
		},
		{
			name: "very small numbers",
			nums: []int{-1000000000, -999999999, -999999998},
			want: 3,
		},
		{
			name: "mixed with zero",
			nums: []int{0, 0, 1, 2},
			want: 3,
		},
		{
			name: "single zero",
			nums: []int{0},
			want: 1,
		},
		{
			name: "complex case",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "duplicates scattered",
			nums: []int{1, 2, 0, 1, 2, 3},
			want: 4,
		},
		{
			name: "long sequence mixed with singles",
			nums: []int{9, 1, 2, 3, 4, 5, 6, 100, 200},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
