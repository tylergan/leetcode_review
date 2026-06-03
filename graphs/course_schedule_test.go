package graphs

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example 1 - one prerequisite",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}},
			want:          true,
		},
		{
			name:          "example 2 - direct cycle",
			numCourses:    2,
			prerequisites: [][]int{{0, 1}, {1, 0}},
			want:          false,
		},
		{
			name:          "no prerequisites",
			numCourses:    4,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "linear dependency chain",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "longer cycle",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			want:          false,
		},
		{
			name:          "disconnected graph with one cycle",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {3, 2}, {2, 3}},
			want:          false,
		},
		{
			name:          "disconnected acyclic graph",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 0}, {4, 3}, {5, 4}},
			want:          true,
		},
		{
			name:          "multiple prerequisites for one course",
			numCourses:    4,
			prerequisites: [][]int{{3, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Fatalf("canFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
