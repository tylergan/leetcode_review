package linked_list

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 3, 2, 2},
			want: 2,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 4, 4},
			want: 4,
		},
		{
			name: "smallest input",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "duplicate appears twice",
			nums: []int{2, 1, 3, 4, 2},
			want: 2,
		},
		{
			name: "duplicate near start",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "duplicate near end",
			nums: []int{1, 4, 6, 2, 5, 3, 6},
			want: 6,
		},
		{
			name: "duplicate appears multiple times",
			nums: []int{5, 4, 3, 2, 1, 5, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			if got != tt.want {
				t.Fatalf("findDuplicate(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
