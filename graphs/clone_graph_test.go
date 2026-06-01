package graphs

import (
	"slices"
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList)+1)
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &Node{Val: i}
	}

	for val, neighbours := range adjList {
		node := nodes[val+1]
		for _, neighbourVal := range neighbours {
			node.Neighbours = append(node.Neighbours, nodes[neighbourVal])
		}
	}

	return nodes[1]
}

func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	seen := map[*Node]bool{node: true}
	queue := []*Node{node}
	maxVal := node.Val
	nodesByVal := map[int]*Node{node.Val: node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		maxVal = max(maxVal, curr.Val)
		nodesByVal[curr.Val] = curr

		for _, neighbour := range curr.Neighbours {
			if !seen[neighbour] {
				seen[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	adjList := make([][]int, maxVal)
	for val := 1; val <= maxVal; val++ {
		node := nodesByVal[val]
		if node == nil {
			continue
		}

		for _, neighbour := range node.Neighbours {
			adjList[val-1] = append(adjList[val-1], neighbour.Val)
		}
		slices.Sort(adjList[val-1])
	}

	return adjList
}

func collectNodes(node *Node) map[int]*Node {
	if node == nil {
		return map[int]*Node{}
	}

	seen := map[*Node]bool{node: true}
	queue := []*Node{node}
	res := map[int]*Node{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		res[curr.Val] = curr

		for _, neighbour := range curr.Neighbours {
			if !seen[neighbour] {
				seen[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	return res
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "example 1 - line graph",
			adjList: [][]int{{2}, {1, 3}, {2}},
		},
		{
			name:    "example 2 - isolated node",
			adjList: [][]int{{}},
		},
		{
			name:    "cycle graph",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "branching graph",
			adjList: [][]int{{2, 3}, {1, 4}, {1, 4}, {2, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			clone := cloneGraph(original)

			got := graphToAdjList(clone)
			for i := range got {
				slices.Sort(got[i])
			}
			for i := range tt.adjList {
				slices.Sort(tt.adjList[i])
			}
			if !slices.EqualFunc(got, tt.adjList, slices.Equal[[]int]) {
				t.Fatalf("cloneGraph(%v) adjacency = %v, want %v", tt.adjList, got, tt.adjList)
			}

			originalNodes := collectNodes(original)
			clonedNodes := collectNodes(clone)
			for val, originalNode := range originalNodes {
				if clonedNodes[val] == originalNode {
					t.Fatalf("node %d was not deep-copied", val)
				}
			}
		})
	}
}

func TestCloneGraphNil(t *testing.T) {
	if got := cloneGraph(nil); got != nil {
		t.Fatalf("cloneGraph(nil) = %v, want nil", got)
	}
}
