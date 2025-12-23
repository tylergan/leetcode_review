package sliding_window

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1: buy low sell high",
			prices: []int{10, 1, 5, 6, 7, 1},
			want:   6,
		},
		{
			name:   "example 2: no profit",
			prices: []int{10, 8, 7, 5, 2},
			want:   0,
		},
		{
			name:   "single day",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "strictly increasing",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "strictly decreasing",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "multiple valleys",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "equal prices",
			prices: []int{3, 3, 3},
			want:   0,
		},
		{
			name:   "late best sell",
			prices: []int{2, 4, 1, 8},
			want:   7,
		},
		{
			name:   "early best sell",
			prices: []int{2, 4, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
