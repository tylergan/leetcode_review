package two_pointers

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example 1: standard case",
			height: []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			want:   9,
		},
		{
			name:   "single element",
			height: []int{5},
			want:   0,
		},
		{
			name:   "two elements - no water",
			height: []int{3, 5},
			want:   0,
		},
		{
			name:   "three elements - valley",
			height: []int{3, 0, 3},
			want:   3,
		},
		{
			name:   "all zeros",
			height: []int{0, 0, 0, 0},
			want:   0,
		},
		{
			name:   "no water - increasing",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "no water - decreasing",
			height: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "simple valley",
			height: []int{5, 1, 5},
			want:   4,
		},
		{
			name:   "multiple valleys",
			height: []int{5, 1, 5, 1, 5},
			want:   8,
		},
		{
			name:   "wide valley",
			height: []int{5, 0, 0, 0, 5},
			want:   15,
		},
		{
			name:   "stepped valley",
			height: []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			want:   16,
		},
		{
			name:   "mountain shape",
			height: []int{0, 1, 2, 3, 2, 1, 0},
			want:   0,
		},
		{
			name:   "flat with walls",
			height: []int{4, 0, 0, 0, 4},
			want:   12,
		},
		{
			name:   "unequal walls",
			height: []int{3, 0, 0, 0, 5},
			want:   9,
		},
		{
			name:   "small wall on left",
			height: []int{2, 0, 5},
			want:   2,
		},
		{
			name:   "small wall on right",
			height: []int{5, 0, 2},
			want:   2,
		},
		{
			name:   "complex pattern",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "stairs up then down",
			height: []int{1, 2, 3, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "barriers with gaps",
			height: []int{5, 2, 1, 2, 1, 5},
			want:   14,
		},
		{
			name:   "highest in middle",
			height: []int{3, 0, 2, 0, 4, 0, 2, 0, 3},
			want:   14,
		},
		{
			name:   "equal heights at ends",
			height: []int{3, 0, 1, 0, 2, 0, 3},
			want:   12,
		},
		{
			name:   "multiple peaks",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "large walls",
			height: []int{1000, 0, 1000},
			want:   1000,
		},
		{
			name:   "large array",
			height: []int{1000, 1, 1, 1, 1, 1, 1, 1, 1, 1000},
			want:   7992,
		},
		{
			name:   "zigzag pattern",
			height: []int{5, 2, 5, 2, 5},
			want:   6,
		},
		{
			name:   "starts and ends with zero",
			height: []int{0, 3, 0, 3, 0},
			want:   3,
		},
		{
			name:   "single dip",
			height: []int{2, 0, 2},
			want:   2,
		},
		{
			name:   "multiple dips",
			height: []int{3, 0, 1, 0, 3},
			want:   8,
		},
		{
			name:   "gradual descent then rise",
			height: []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			want:   16,
		},
		{
			name:   "plateau in valley",
			height: []int{5, 1, 1, 1, 5},
			want:   12,
		},
		{
			name:   "narrow gaps",
			height: []int{4, 1, 4, 1, 4},
			want:   6,
		},
		{
			name:   "wide gaps",
			height: []int{5, 0, 0, 0, 0, 0, 5},
			want:   25,
		},
		{
			name:   "asymmetric valley",
			height: []int{6, 0, 0, 0, 3},
			want:   9,
		},
		{
			name:   "step pattern",
			height: []int{3, 0, 2, 0, 4},
			want:   7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrapProperties(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		validate func(t *testing.T, height []int, result int)
	}{
		{
			name:   "result is non-negative",
			height: []int{1, 2, 3, 4},
			validate: func(t *testing.T, height []int, result int) {
				if result < 0 {
					t.Errorf("trap() returned negative value: %d", result)
				}
			},
		},
		{
			name:   "single element traps no water",
			height: []int{10},
			validate: func(t *testing.T, height []int, result int) {
				if result != 0 {
					t.Errorf("trap() with single element = %d, want 0", result)
				}
			},
		},
		{
			name:   "all zeros traps no water",
			height: []int{0, 0, 0, 0},
			validate: func(t *testing.T, height []int, result int) {
				if result != 0 {
					t.Errorf("trap() with all zeros = %d, want 0", result)
				}
			},
		},
		{
			name:   "increasing sequence traps no water",
			height: []int{1, 2, 3, 4, 5},
			validate: func(t *testing.T, height []int, result int) {
				if result != 0 {
					t.Errorf("trap() with increasing sequence = %d, want 0", result)
				}
			},
		},
		{
			name:   "decreasing sequence traps no water",
			height: []int{5, 4, 3, 2, 1},
			validate: func(t *testing.T, height []int, result int) {
				if result != 0 {
					t.Errorf("trap() with decreasing sequence = %d, want 0", result)
				}
			},
		},
		{
			name:   "result bounded by array size and max height",
			height: []int{10, 0, 0, 0, 10},
			validate: func(t *testing.T, height []int, result int) {
				maxHeight := 0
				for _, h := range height {
					if h > maxHeight {
						maxHeight = h
					}
				}
				maxPossible := maxHeight * len(height)
				if result > maxPossible {
					t.Errorf("trap() = %d exceeds maximum possible %d", result, maxPossible)
				}
			},
		},
		{
			name:   "simple valley calculation",
			height: []int{5, 0, 5},
			validate: func(t *testing.T, height []int, result int) {
				expected := 5 // One unit of width at height 5
				if result != expected {
					t.Errorf("trap() = %d, want %d", result, expected)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trap(tt.height)
			tt.validate(t, tt.height, result)
		})
	}
}
