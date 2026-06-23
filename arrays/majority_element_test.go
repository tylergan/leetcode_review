package arrays

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1 - majority not contiguous",
			nums: []int{5, 5, 1, 1, 1, 5, 5},
			want: 5,
		},
		{
			name: "example 2 - all same",
			nums: []int{2, 2, 2},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{7},
			want: 7,
		},
		{
			name: "majority at the front",
			nums: []int{4, 4, 4, 1, 2},
			want: 4,
		},
		{
			name: "majority at the back",
			nums: []int{1, 2, 3, 3, 3, 3, 3},
			want: 3,
		},
		{
			name: "candidate gets replaced before the true majority wins",
			nums: []int{8, 9, 8, 9, 8},
			want: 8,
		},
		{
			name: "negative majority element",
			nums: []int{-1, -1, -1, 5, 5},
			want: -1,
		},
		{
			name: "majority exactly fills half plus one",
			nums: []int{6, 1, 6, 2, 6, 3, 6},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Fatalf("majorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
