package binary_search

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		//{
		//	name:  "example 1",
		//	piles: []int{1, 4, 3, 2},
		//	h:     9,
		//	want:  2,
		//},
		//{
		//	name:  "example 2",
		//	piles: []int{25, 10, 23, 4},
		//	h:     4,
		//	want:  25,
		//},
		//{
		//	name:  "single pile exact hours",
		//	piles: []int{8},
		//	h:     4,
		//	want:  2,
		//},
		//{
		//	name:  "single pile one hour",
		//	piles: []int{8},
		//	h:     1,
		//	want:  8,
		//},
		//{
		//	name:  "h equals number of piles",
		//	piles: []int{3, 6, 7, 11},
		//	h:     4,
		//	want:  11,
		//},
		{
			name:  "h equals total bananas",
			piles: []int{3, 6, 7, 11},
			h:     27,
			want:  1,
		},
		//{
		//	name:  "target requires ceiling per pile",
		//	piles: []int{30, 11, 23, 4, 20},
		//	h:     5,
		//	want:  30,
		//},
		//{
		//	name:  "more hours allows lower speed",
		//	piles: []int{30, 11, 23, 4, 20},
		//	h:     6,
		//	want:  23,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			if got != tt.want {
				t.Fatalf("minEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.want)
			}
		})
	}
}
