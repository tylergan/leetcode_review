package linked_list

/*
You are given the head of a singly linked list head and a positive integer k.

You must reverse the first k nodes in the linked list, and then reverse the next k nodes, and so on. If there are fewer than k nodes left, leave the nodes as they are.

Return the modified list after reversing the nodes in each group of k.

You are only allowed to modify the nodes' next pointers, not the values of the nodes.

Example 1:



Input: head = [1,2,3,4,5,6], k = 3

Output: [3,2,1,6,5,4]
Example 2:



Input: head = [1,2,3,4,5], k = 3

Output: [3,2,1,4,5]
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	dummy := &ListNode{}
	curr := dummy
	for p != nil {
		q := p
		i := 0
		for q != nil && i < k { // check if there are "k" nodes
			q = q.Next
			i++
		}

		if i < k { // if we have not enough for reversing "k", just add it to our curr and exit
			curr.Next = p
			break
		}

		// Otherwise, begin reversing our group of "k"
		i = 0
		prev := q      // our end should reference where "k" terminates (the node after our group of "k")
		groupHead := p // keep track of our previously old groupHead (soon to become tail of the group)
		for p != nil && i < k {
			tmp := p.Next
			p.Next = prev
			prev = p
			p = tmp
			i++
		}
		curr.Next = prev // prev now the head of our reversed group of "k"
		curr = groupHead // move forward to the end of the group

	}

	return dummy.Next
}
