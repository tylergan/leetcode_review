package heappq

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "example 1",
			stones: []int{2, 3, 6, 2, 4},
			want:   1,
		},
		{
			name:   "example 2",
			stones: []int{1, 2},
			want:   1,
		},
		{
			name:   "single stone",
			stones: []int{5},
			want:   5,
		},
		{
			name:   "all equal - even count",
			stones: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "all equal - odd count",
			stones: []int{3, 3, 3},
			want:   3,
		},
		{
			name:   "two equal stones",
			stones: []int{4, 4},
			want:   0,
		},
		{
			name:   "large difference",
			stones: []int{1, 100},
			want:   99,
		},
		{
			name:   "descending order",
			stones: []int{10, 7, 3, 1},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.stones); got != tt.want {
				t.Fatalf("lastStoneWeight(%v) = %d, want %d", tt.stones, got, tt.want)
			}
		})
	}
}
