package two_pointers

import (
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		// want holds the expected remaining elements (order independent), so
		// its length is the expected k. It mirrors the LeetCode custom judge.
		want []int
	}{
		{
			name: "example 1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: []int{2, 2},
		},
		{
			name: "example 2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: []int{0, 0, 1, 3, 4},
		},
		{
			name: "empty array",
			nums: []int{},
			val:  5,
			want: []int{},
		},
		{
			name: "val not present",
			nums: []int{1, 2, 3},
			val:  5,
			want: []int{1, 2, 3},
		},
		{
			name: "all elements equal val",
			nums: []int{2, 2, 2, 2},
			val:  2,
			want: []int{},
		},
		{
			name: "single element equal val",
			nums: []int{3},
			val:  3,
			want: []int{},
		},
		{
			name: "single element not equal val",
			nums: []int{5},
			val:  3,
			want: []int{5},
		},
		{
			name: "val at both ends",
			nums: []int{4, 1, 2, 3, 4},
			val:  4,
			want: []int{1, 2, 3},
		},
		{
			name: "consecutive vals in the middle",
			nums: []int{1, 4, 4, 4, 2},
			val:  4,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			k := removeElement(nums, tt.val)

			if k != len(tt.want) {
				t.Fatalf("removeElement(%v, %d) = %d, want %d", tt.nums, tt.val, k, len(tt.want))
			}

			got := append([]int(nil), nums[:k]...)
			sort.Ints(got)
			want := append([]int(nil), tt.want...)
			sort.Ints(want)

			for idx := range want {
				if got[idx] != want[idx] {
					t.Fatalf("removeElement(%v, %d) first %d elements = %v, want %v (order independent)", tt.nums, tt.val, k, got, want)
				}
				if got[idx] == tt.val {
					t.Fatalf("removeElement(%v, %d) left val %d within the first %d elements: %v", tt.nums, tt.val, tt.val, k, got)
				}
			}
		})
	}
}
