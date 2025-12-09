package two_pointers

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "example 1: mixed heights",
			heights: []int{1, 7, 2, 5, 4, 7, 3, 6},
			want:    36,
		},
		{
			name:    "example 2: all same heights",
			heights: []int{2, 2, 2},
			want:    4,
		},
		{
			name:    "minimum size - two elements",
			heights: []int{1, 1},
			want:    1,
		},
		{
			name:    "minimum size - different heights",
			heights: []int{5, 10},
			want:    5,
		},
		{
			name:    "increasing heights",
			heights: []int{1, 2, 3, 4, 5},
			want:    6,
		},
		{
			name:    "decreasing heights",
			heights: []int{5, 4, 3, 2, 1},
			want:    6,
		},
		{
			name:    "maximum at ends",
			heights: []int{10, 1, 1, 1, 10},
			want:    40,
		},
		{
			name:    "maximum in middle",
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:    49,
		},
		{
			name:    "tall bars at edges",
			heights: []int{100, 1, 2, 3, 100},
			want:    400,
		},
		{
			name:    "with zeros at start",
			heights: []int{0, 0, 5, 10},
			want:    5,
		},
		{
			name:    "with zeros at end",
			heights: []int{10, 5, 0, 0},
			want:    5,
		},
		{
			name:    "with zeros in middle",
			heights: []int{5, 0, 0, 0, 5},
			want:    20,
		},
		{
			name:    "all zeros",
			heights: []int{0, 0, 0},
			want:    0,
		},
		{
			name:    "one tall one short",
			heights: []int{1, 100},
			want:    1,
		},
		{
			name:    "symmetric pattern",
			heights: []int{1, 2, 3, 4, 3, 2, 1},
			want:    8,
		},
		{
			name:    "large numbers",
			heights: []int{1000, 1000},
			want:    1000,
		},
		{
			name:    "large array with max at edges",
			heights: []int{1000, 1, 1, 1, 1, 1, 1, 1, 1, 1000},
			want:    9000,
		},
		{
			name:    "valley shape",
			heights: []int{5, 1, 1, 1, 5},
			want:    20,
		},
		{
			name:    "peak shape",
			heights: []int{1, 5, 5, 5, 1},
			want:    10,
		},
		{
			name:    "alternating high low",
			heights: []int{10, 1, 10, 1, 10},
			want:    40,
		},
		{
			name:    "gradually increasing then peak",
			heights: []int{1, 2, 4, 8, 16, 8, 4, 2, 1},
			want:    16,
		},
		{
			name:    "single tall bar among short ones",
			heights: []int{1, 1, 1, 100, 1, 1, 1},
			want:    6,
		},
		{
			name:    "two tall bars far apart",
			heights: []int{50, 1, 1, 1, 1, 1, 1, 1, 50},
			want:    400,
		},
		{
			name:    "sequential same heights",
			heights: []int{5, 5, 5, 5, 5},
			want:    20,
		},
		{
			name:    "close tall bars",
			heights: []int{1, 100, 100, 1},
			want:    100,
		},
		{
			name:    "barrier in middle",
			heights: []int{10, 5, 8, 5, 10},
			want:    40,
		},
		{
			name:    "wide container with small height",
			heights: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			want:    9,
		},
		{
			name:    "narrow container with large height",
			heights: []int{100, 100},
			want:    100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.heights)
			if got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAreaProperties(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		validate func(t *testing.T, heights []int, result int)
	}{
		{
			name:    "result is non-negative",
			heights: []int{1, 2, 3, 4},
			validate: func(t *testing.T, heights []int, result int) {
				if result < 0 {
					t.Errorf("maxArea() returned negative value: %d", result)
				}
			},
		},
		{
			name:    "result is at most max_height * (n-1)",
			heights: []int{5, 10, 3, 8},
			validate: func(t *testing.T, heights []int, result int) {
				maxHeight := 0
				for _, h := range heights {
					if h > maxHeight {
						maxHeight = h
					}
				}
				maxPossible := maxHeight * (len(heights) - 1)
				if result > maxPossible {
					t.Errorf("maxArea() = %d exceeds maximum possible %d", result, maxPossible)
				}
			},
		},
		{
			name:    "all zeros gives zero area",
			heights: []int{0, 0, 0, 0},
			validate: func(t *testing.T, heights []int, result int) {
				if result != 0 {
					t.Errorf("maxArea() with all zeros = %d, want 0", result)
				}
			},
		},
		{
			name:    "minimum size array",
			heights: []int{3, 7},
			validate: func(t *testing.T, heights []int, result int) {
				expected := min(heights[0], heights[1]) * 1
				if result != expected {
					t.Errorf("maxArea() = %d, want %d", result, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxArea(tt.heights)
			tt.validate(t, tt.heights, result)
		})
	}
}
