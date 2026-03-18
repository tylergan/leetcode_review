package heappq

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name    string
		adds    []int
		expects []float64
	}{
		{
			name:    "example from problem",
			adds:    []int{1, 3, 2},
			expects: []float64{1.0, 2.0, 2.0},
		},
		{
			name:    "single element",
			adds:    []int{5},
			expects: []float64{5.0},
		},
		{
			name:    "two elements",
			adds:    []int{1, 2},
			expects: []float64{1.0, 1.5},
		},
		{
			name:    "descending order",
			adds:    []int{3, 2, 1},
			expects: []float64{3.0, 2.5, 2.0},
		},
		{
			name:    "ascending order",
			adds:    []int{1, 2, 3},
			expects: []float64{1.0, 1.5, 2.0},
		},
		{
			name:    "duplicates",
			adds:    []int{5, 5, 5, 5},
			expects: []float64{5.0, 5.0, 5.0, 5.0},
		},
		{
			name:    "negative values",
			adds:    []int{-3, -1, -2},
			expects: []float64{-3.0, -2.0, -2.0},
		},
		{
			name:    "mixed positive and negative",
			adds:    []int{-1, 0, 1},
			expects: []float64{-1.0, -0.5, 0.0},
		},
		{
			name:    "even count median is average of two middle",
			adds:    []int{1, 2, 3, 4},
			expects: []float64{1.0, 1.5, 2.0, 2.5},
		},
		{
			name:    "large spread",
			adds:    []int{1, 5, 10},
			expects: []float64{1.0, 3.0, 5.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := NewMedianFinder()
			for i, num := range tt.adds {
				mf.AddNum(num)
				got := mf.FindMedian()
				if math.Abs(got-tt.expects[i]) > 1e-5 {
					t.Fatalf("after adding %d (step %d): FindMedian() = %v, want %v", num, i, got, tt.expects[i])
				}
			}
		})
	}
}
