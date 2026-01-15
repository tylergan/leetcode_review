package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		cycleIndex int
		want       bool
	}{
		{
			name:       "single node no cycle",
			values:     []int{1},
			cycleIndex: -1,
			want:       false,
		},
		{
			name:       "single node self cycle",
			values:     []int{1},
			cycleIndex: 0,
			want:       true,
		},
		{
			name:       "two nodes no cycle",
			values:     []int{1, 2},
			cycleIndex: -1,
			want:       false,
		},
		{
			name:       "two nodes cycle to head",
			values:     []int{1, 2},
			cycleIndex: 0,
			want:       true,
		},
		{
			name:       "longer list cycle in middle",
			values:     []int{1, 2, 3, 4, 5},
			cycleIndex: 2,
			want:       true,
		},
		{
			name:       "longer list no cycle",
			values:     []int{1, 2, 3, 4, 5},
			cycleIndex: -1,
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := listWithCycle(tt.values, tt.cycleIndex)
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
