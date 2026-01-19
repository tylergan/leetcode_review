package linked_list

import "testing"

type randomListCase struct {
	name      string
	values    []int
	randomIdx []int
}

func buildRandomList(values []int, randomIdx []int) (*Node, []*Node) {
	if len(values) == 0 {
		return nil, nil
	}
	nodes := make([]*Node, len(values))
	for i, v := range values {
		nodes[i] = &Node{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	for i, idx := range randomIdx {
		if idx >= 0 {
			nodes[i].Random = nodes[idx]
		}
	}
	return nodes[0], nodes
}

func collectNodes(head *Node) []*Node {
	var nodes []*Node
	for curr := head; curr != nil; curr = curr.Next {
		nodes = append(nodes, curr)
	}
	return nodes
}

func TestCopyRandomList(t *testing.T) {
	tests := []randomListCase{
		{
			name:      "empty list",
			values:    nil,
			randomIdx: nil,
		},
		{
			name:      "single node nil random",
			values:    []int{7},
			randomIdx: []int{-1},
		},
		{
			name:      "single node random to self",
			values:    []int{9},
			randomIdx: []int{0},
		},
		{
			name:      "multiple nodes mixed randoms",
			values:    []int{3, 7, 4, 5},
			randomIdx: []int{-1, 3, 0, 1},
		},
		{
			name:      "randoms to same node",
			values:    []int{1, 2, 3},
			randomIdx: []int{-1, 2, 2},
		},
		{
			name:      "repeated values",
			values:    []int{5, 5, 5},
			randomIdx: []int{2, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, origNodes := buildRandomList(tt.values, tt.randomIdx)

			origNext := make([]*Node, len(origNodes))
			origRandom := make([]*Node, len(origNodes))
			for i, node := range origNodes {
				origNext[i] = node.Next
				origRandom[i] = node.Random
			}

			gotHead := copyRandomList(head)

			if head == nil {
				if gotHead != nil {
					t.Fatalf("copyRandomList() = %v, want nil", gotHead)
				}
				return
			}

			copyNodes := collectNodes(gotHead)
			if len(copyNodes) != len(origNodes) {
				t.Fatalf("copyRandomList() len = %d, want %d", len(copyNodes), len(origNodes))
			}

			origIndex := make(map[*Node]int, len(origNodes))
			for i, node := range origNodes {
				origIndex[node] = i
			}
			copyIndex := make(map[*Node]int, len(copyNodes))
			for i, node := range copyNodes {
				copyIndex[node] = i
			}

			for i, orig := range origNodes {
				copyNode := copyNodes[i]
				if orig == copyNode {
					t.Fatalf("node %d was not copied", i)
				}
				if orig.Val != copyNode.Val {
					t.Fatalf("node %d val = %d, want %d", i, copyNode.Val, orig.Val)
				}

				if orig.Next == nil {
					if copyNode.Next != nil {
						t.Fatalf("node %d next = %v, want nil", i, copyNode.Next)
					}
				} else {
					if copyNode.Next != copyNodes[i+1] {
						t.Fatalf("node %d next mismatch", i)
					}
				}

				if orig.Random == nil {
					if copyNode.Random != nil {
						t.Fatalf("node %d random = %v, want nil", i, copyNode.Random)
					}
				} else {
					randomIdx := origIndex[orig.Random]
					if copyNode.Random != copyNodes[randomIdx] {
						t.Fatalf("node %d random points to %d, want %d", i, copyIndex[copyNode.Random], randomIdx)
					}
					if _, ok := origIndex[copyNode.Random]; ok {
						t.Fatalf("node %d random points to original list", i)
					}
				}
			}

			for i, node := range origNodes {
				if node.Next != origNext[i] {
					t.Fatalf("original node %d next pointer changed", i)
				}
				if node.Random != origRandom[i] {
					t.Fatalf("original node %d random pointer changed", i)
				}
			}
		})
	}
}
