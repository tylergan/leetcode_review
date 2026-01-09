package binary_search

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 5, 6, 1, 2},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 5, 0, 1, 2, 3},
			want: 0,
		},
		{
			name: "example 3: not rotated",
			nums: []int{4, 5, 6, 7},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{10},
			want: 10,
		},
		{
			name: "two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "two elements not rotated",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "rotation by one",
			nums: []int{6, 1, 2, 3, 4, 5},
			want: 1,
		},
		{
			name: "rotation in middle",
			nums: []int{30, 40, 50, 5, 10, 20},
			want: 5,
		},
		{
			name: "negative values",
			nums: []int{0, 3, 5, -10, -5, -2},
			want: -10,
		},
		{
			name: "minimum at end",
			nums: []int{2, 3, 4, 5, 6, 1},
			want: 1,
		},
		{
			name: "minimum at start",
			nums: []int{-3, -2, -1, 0, 1},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Fatalf("findMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
