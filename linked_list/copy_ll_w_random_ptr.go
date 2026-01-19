package linked_list

/*
You are given the head of a linked list of length n. Unlike a singly linked list, each node contains an additional pointer random, which may point to any node in the list, or null.

Create a deep copy of the list.

The deep copy should consist of exactly n new nodes, each including:

The original value val of the copied node
A next pointer to the new node corresponding to the next pointer of the original node
A random pointer to the new node corresponding to the random pointer of the original node
Note: None of the pointers in the new list should point to nodes in the original list.

Return the head of the copied linked list.

In the examples, the linked list is represented as a list of n nodes. Each node is represented as a pair of [val, random_index] where random_index is the index of the node (0-indexed) that the random pointer points to, or null if it does not point to any node.

Example 1:



Input: head = [[3,null],[7,3],[4,0],[5,1]]

Output: [[3,null],[7,3],[4,0],[5,1]]
Example 2:



Input: head = [[1,null],[2,2],[3,2]]

Output: [[1,null],[2,2],[3,2]]
*/

func copyRandomList(head *Node) *Node {
	mapping := map[*Node]*Node{}
	dummy := &Node{}
	prev := dummy
	p := head
	// create a mapping of original to copies
	for p != nil {
		copied := &Node{Val: p.Val}
		prev.Next = copied
		mapping[p] = copied // map original to copies

		// move forward
		prev.Next = copied
		prev = prev.Next
		p = p.Next
	}
	// now, assign randoms to copied
	p, prev = head, dummy.Next
	for p != nil {
		if p.Random != nil {
			copied, _ := mapping[p.Random]
			prev.Random = copied
		}
		p = p.Next
		prev = prev.Next
	}
	return dummy.Next
}
