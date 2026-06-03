package graphs

/*
Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

Example 1:

Input:
n = 5
edges = [[0, 1], [0, 2], [0, 3], [1, 4]]

Output:
true
Example 2:

Input:
n = 5
edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]

Output:
false
Note:

You can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.
Constraints:

1 <= n <= 100
0 <= edges.length <= n * (n - 1) / 2
*/

type UnionFind struct {
	parent      []int
	size        []int
	nComponents int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size, nComponents: n}
}

func (u *UnionFind) Union(x, y int) bool {
	a, b := u.find(x), u.find(y)
	if a == b {
		return false
	}

	if u.size[a] < u.size[b] {
		a, b = b, a
	}

	u.parent[b] = a
	u.size[a] += u.size[b]
	u.nComponents--

	return true
}

func (u *UnionFind) find(x int) int {
	if x == u.parent[x] {
		return x
	}
	u.parent[x] = u.find(u.parent[x])
	return u.parent[x]
}

func validTree(n int, edges [][]int) bool {
	uf := NewUnionFind(n)
	for _, edge := range edges {
		if !uf.Union(edge[0], edge[1]) { // Check for cycle
			return false
		}
	}

	return uf.nComponents == 1
}

func validTreeDFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}

	adjList := map[int][]int{}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	seen := map[int]bool{}

	var dfs func(int)
	dfs = func(node int) {
		seen[node] = true

		for _, neighbour := range adjList[node] {
			if seen[neighbour] {
				continue
			}

			dfs(neighbour)
		}
	}

	dfs(0)

	return len(seen) == n
}
