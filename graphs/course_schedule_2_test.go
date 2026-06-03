package graphs

import "testing"

func isValidCourseOrder(numCourses int, prerequisites [][]int, order []int) bool {
	if len(order) != numCourses {
		return false
	}

	positionByCourse := map[int]int{}
	for i, course := range order {
		if course < 0 || course >= numCourses {
			return false
		}
		if _, ok := positionByCourse[course]; ok {
			return false
		}
		positionByCourse[course] = i
	}

	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		if positionByCourse[prerequisite] > positionByCourse[course] {
			return false
		}
	}

	return true
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		wantPossible  bool
	}{
		{
			name:          "example 1 - one prerequisite",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}},
			wantPossible:  true,
		},
		{
			name:          "example 2 - cycle returns empty",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			wantPossible:  false,
		},
		{
			name:          "no prerequisites returns all courses",
			numCourses:    4,
			prerequisites: [][]int{},
			wantPossible:  true,
		},
		{
			name:          "linear dependency chain",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			wantPossible:  true,
		},
		{
			name:          "multiple prerequisites for one course",
			numCourses:    4,
			prerequisites: [][]int{{3, 0}, {3, 1}, {3, 2}},
			wantPossible:  true,
		},
		{
			name:          "disconnected acyclic components",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 0}, {4, 3}, {5, 4}},
			wantPossible:  true,
		},
		{
			name:          "disconnected graph with one cycle returns empty",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {3, 2}, {2, 3}},
			wantPossible:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)

			if !tt.wantPossible {
				if len(got) != 0 {
					t.Fatalf("findOrder(%d, %v) = %v, want empty order", tt.numCourses, tt.prerequisites, got)
				}
				return
			}

			if !isValidCourseOrder(tt.numCourses, tt.prerequisites, got) {
				t.Fatalf("findOrder(%d, %v) = %v, want valid course order", tt.numCourses, tt.prerequisites, got)
			}
		})
	}
}
