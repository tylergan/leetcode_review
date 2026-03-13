package heappq

import "testing"

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name    string
		k       int
		nums    []int
		adds    []int
		expects []int
	}{
		{
			name:    "example 1",
			k:       3,
			nums:    []int{1, 2, 3, 3},
			adds:    []int{3, 5, 6, 7, 8},
			expects: []int{3, 3, 3, 5, 6},
		},
		{
			name:    "negative values",
			k:       2,
			nums:    []int{-4, -5},
			adds:    []int{-2, -3, -1},
			expects: []int{-4, -3, -2},
		},
		{
			name:    "duplicates",
			k:       2,
			nums:    []int{5, 5, 5},
			adds:    []int{5, 5, 10},
			expects: []int{5, 5, 5},
		},
		{
			name:    "single initial element",
			k:       1,
			nums:    []int{10},
			adds:    []int{5, 15, 20},
			expects: []int{10, 15, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kl := Constructor(tt.k, tt.nums)
			for i, val := range tt.adds {
				got := kl.Add(val)
				if got != tt.expects[i] {
					t.Fatalf("Add(%d) = %d, want %d (step %d)", val, got, tt.expects[i], i)
				}
			}
		})
	}
}
