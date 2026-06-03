package graphs

/*
You are given an array prerequisites where prerequisites[i] = [a, b] indicates that you must take course b first if you want to take course a.

The pair [0, 1], indicates that must take course 1 before taking course 0.

There are a total of numCourses courses you are required to take, labeled from 0 to numCourses - 1.

Return true if it is possible to finish all courses, otherwise return false.

Example 1:

Input: numCourses = 2, prerequisites = [[0,1]]

Output: true
Explanation: First take course 1 (no prerequisites) and then take course 0.

Example 2:

Input: numCourses = 2, prerequisites = [[0,1],[1,0]]

Output: false
Explanation: In order to take course 1 you must take course 0, and to take course 0 you must take course 1. So it is impossible.

Constraints:

1 <= numCourses <= 1000
0 <= prerequisites.length <= 1000
prerequisites[i].length == 2
0 <= a[i], b[i] < numCourses
All prerequisite pairs are unique.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Kahn's algorithm.
	adjList := map[int][]int{}
	indegrees := make([]int, numCourses)

	for _, prereq := range prerequisites { // Calculate indegrees and build graph.
		to, from := prereq[0], prereq[1]
		adjList[from] = append(adjList[from], to)
		indegrees[to]++
	}

	q := []int{}
	for i, indeg := range indegrees { // Add all source nodes of graph (indegree = 0).
		if indeg == 0 {
			q = append(q, i)
		}
	}

	processed := 0 // See if we can process all nodes.
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		processed++

		neighbours := adjList[curr]
		for _, neigh := range neighbours {
			indegrees[neigh]--
			if indegrees[neigh] == 0 {
				q = append(q, neigh)
			}
		}
	}

	return processed == numCourses
}
