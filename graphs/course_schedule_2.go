package graphs

/*
You are given an array prerequisites where prerequisites[i] = [a, b] indicates that you must take course b first if you want to take course a.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
There are a total of numCourses courses you are required to take, labeled from 0 to numCourses - 1.

Return a valid ordering of courses you can take to finish all courses. If there are many valid answers, return any of them. If it's not possible to finish all courses, return an empty array.

Example 1:

Input: numCourses = 3, prerequisites = [[1,0]]

Output: [0,1,2]
Explanation: We must ensure that course 0 is taken before course 1.

Example 2:

Input: numCourses = 3, prerequisites = [[0,1],[1,2],[2,0]]

Output: []
Explanation: It's impossible to finish all courses.

Constraints:

1 <= numCourses <= 1000
0 <= prerequisites.length <= 1000
All prerequisite pairs are unique.
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := map[int][]int{}
	indegrees := make([]int, numCourses)

	for _, prereq := range prerequisites {
		to, from := prereq[0], prereq[1]
		adjList[from] = append(adjList[from], to)
		indegrees[to]++
	}

	q := []int{}
	for i, indeg := range indegrees {
		if indeg == 0 {
			q = append(q, i)
		}
	}

	res := []int{}
	processed := 0

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		processed++
		res = append(res, curr)

		neighbours := adjList[curr]
		for _, neigh := range neighbours {
			indegrees[neigh]--
			if indegrees[neigh] == 0 {
				q = append(q, neigh)
			}
		}
	}

	if processed != numCourses {
		return []int{}
	}
	return res
}
