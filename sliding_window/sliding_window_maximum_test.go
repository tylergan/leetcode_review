package sliding_window

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "neetcode example",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "simple example from prompt",
			nums: []int{1, 2, 1, 0, 4, 2, 6},
			k:    3,
			want: []int{2, 2, 4, 4, 6},
		},
		{
			name: "window size one",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "window size equals length",
			nums: []int{9, 11},
			k:    2,
			want: []int{11},
		},
		{
			name: "two elements descending",
			nums: []int{4, -2},
			k:    2,
			want: []int{4},
		},
		{
			name: "all same values",
			nums: []int{2, 2, 2, 2},
			k:    2,
			want: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
