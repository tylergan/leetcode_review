package arrays

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1 - duplicates",
			nums: []int{10, 9, 1, 1, 1, 2, 3, 1},
			want: []int{1, 1, 1, 1, 2, 3, 9, 10},
		},
		{
			name: "example 2",
			nums: []int{5, 10, 2, 1, 3},
			want: []int{1, 2, 3, 5, 10},
		},
		{
			name: "single element",
			nums: []int{7},
			want: []int{7},
		},
		{
			name: "two elements unsorted",
			nums: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "already sorted",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse sorted",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all equal",
			nums: []int{4, 4, 4, 4},
			want: []int{4, 4, 4, 4},
		},
		{
			name: "negative numbers",
			nums: []int{-5, -1, -50000, 0, 3},
			want: []int{-50000, -5, -1, 0, 3},
		},
		{
			name: "boundary values",
			nums: []int{50000, -50000, 0, 50000, -50000},
			want: []int{-50000, -50000, 0, 50000, 50000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArray(tt.nums)
			if len(got) != len(tt.want) {
				t.Fatalf("sortArray() length = %d, want %d", len(got), len(tt.want))
			}
			for idx := range tt.want {
				if got[idx] != tt.want[idx] {
					t.Fatalf("sortArray() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

// Cross-check against the standard library over many random inputs.
func TestSortArrayMatchesStdlib(t *testing.T) {
	rng := rand.New(rand.NewSource(123))

	for iter := 0; iter < 500; iter++ {
		n := rng.Intn(300)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(100001) - 50000 // [-50000, 50000]
		}

		want := append([]int(nil), nums...)
		sort.Ints(want)

		got := sortArray(append([]int(nil), nums...))

		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("iter %d: sortArray(%v) = %v, want %v", iter, nums, got, want)
			}
		}
	}
}
