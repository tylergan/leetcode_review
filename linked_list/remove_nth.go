package linked_list

/*
You are given the beginning of a linked list head, and an integer n.

Remove the nth node from the end of the list and return the beginning of the list.

Example 1:

Input: head = [1,2,3,4], n = 2

Output: [1,2,4]
Example 2:

Input: head = [5], n = 1

Output: []
Example 3:

Input: head = [1,2], n = 2

Output: [2]
Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, head

	// Advance by "n"
	for n > 0 {
		right = right.Next
		n--
	}

	// Find the "n" from end by advancing both left and right until right = nil
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// Now left is at point before target since left started at dummy node
	left.Next = left.Next.Next
	return dummy.Next
}
