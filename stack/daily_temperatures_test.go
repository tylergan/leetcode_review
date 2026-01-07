package stack

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "example 1",
			temperatures: []int{30, 38, 30, 36, 35, 40, 28},
			want:         []int{1, 4, 1, 2, 1, 0, 0},
		},
		{
			name:         "example 2",
			temperatures: []int{22, 21, 20},
			want:         []int{0, 0, 0},
		},
		{
			name:         "strictly increasing",
			temperatures: []int{10, 20, 30, 40},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "strictly decreasing",
			temperatures: []int{40, 30, 20, 10},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "plateau then warmer",
			temperatures: []int{70, 70, 70, 71},
			want:         []int{3, 2, 1, 0},
		},
		{
			name:         "zigzag",
			temperatures: []int{73, 71, 74, 72, 75},
			want:         []int{2, 1, 2, 1, 0},
		},
		{
			name:         "single day",
			temperatures: []int{55},
			want:         []int{0},
		},
		{
			name:         "warm day at end",
			temperatures: []int{60, 50, 40, 70},
			want:         []int{3, 2, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
