package graphs

/*
You are given a connected undirected graph with n nodes labeled from 1 to n. Initially, it contained no cycles and consisted of n-1 edges.

We have now added one additional edge to the graph. The edge has two different vertices chosen from 1 to n, and was not an edge that previously existed in the graph.

The graph is represented as an array edges of length n where edges[i] = [a_i, b_i] represents an edge between nodes a_i and b_i in the graph.

Return an edge that can be removed so that the graph is still a connected non-cyclical graph. If there are multiple answers, return the edge that appears last in the input edges.

Example 1:

Input: edges = [[1,2],[1,3],[3,4],[2,4]]

Output: [2,4]
Example 2:

Input: edges = [[1,2],[1,3],[1,4],[3,4],[4,5]]

Output: [3,4]
Constraints:

n == edges.length
3 <= n <= 100
1 <= edges[i][0] < edges[i][1] <= edges.length
There are no repeated edges and no self-loops in the input.
*/

type RedundantUnionFind struct {
	parent map[int]int
	size   map[int]int
}

func NewRedundantUnionFind() RedundantUnionFind {
	return RedundantUnionFind{
		parent: map[int]int{},
		size:   map[int]int{},
	}
}

func (u *RedundantUnionFind) Union(x, y int) bool {
	a, b := u.find(x), u.find(y)
	if a == b {
		return false // Path already exists; new x-y connection forms a cycle.
	}

	if u.size[a] < u.size[b] {
		a, b = b, a
	}

	u.parent[b] = a
	u.size[a] += u.size[b]

	return true
}

func (u *RedundantUnionFind) find(x int) int {
	if _, ok := u.parent[x]; !ok {
		u.parent[x] = x
		u.size[x] = 1
	}

	if x == u.parent[x] {
		return x
	}

	u.parent[x] = u.find(u.parent[x])
	return u.parent[x]
}

func findRedundantConnection(edges [][]int) []int {
	// Remove the latest edge that forms a cyclic component of the graph. Any removal
	// of that edge should not cause the graph to be disconnected.
	var res []int

	uf := NewRedundantUnionFind()
	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) {
			res = edge
		}
	}

	return res
}
