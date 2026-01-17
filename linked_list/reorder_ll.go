package linked_list

/*
You are given the head of a singly linked-list.

The positions of a linked list of length = 7 for example, can intially be represented as:

[0, 1, 2, 3, 4, 5, 6]

Reorder the nodes of the linked list to be in the following order:

[0, 6, 1, 5, 2, 4, 3]

Notice that in the general case for a list of length = n the nodes are reordered to be in the following order:

[0, n-1, 1, n-2, 2, n-3, ...]

You may not modify the values in the list's nodes, but instead you must reorder the nodes themselves.

Example 1:

Input: head = [2,4,6,8]

Output: [2,8,4,6]
Example 2:

Input: head = [2,4,6,8,10]

Output: [2,10,4,8,6]
Constraints:

1 <= Length of the list <= 1000.
1 <= Node.val <= 1000
*/

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// find the halfway point using slow and fast pointers
	tortoise, hare := head, head.Next
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	// reverse the second half of the list
	var prev *ListNode
	second := tortoise.Next
	tortoise.Next = nil
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// start connecting the values
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}
