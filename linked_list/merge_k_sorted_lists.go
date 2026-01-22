package linked_list

import "container/heap"

/*
You are given an array of k linked lists lists, where each list is sorted in ascending order.

Return the sorted linked list that is the result of merging all of the individual linked lists.

Example 1:

Input: lists = [[1,2,4],[1,3,5],[3,6]]

Output: [1,1,2,3,3,4,5,6]
Example 2:

Input: lists = []

Output: []
Example 3:

Input: lists = [[]]

Output: []
Constraints:

0 <= lists.length <= 1000
0 <= lists[i].length <= 100
-1000 <= lists[i][j] <= 1000
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	mergeTwoListsHelper := func(list1 *ListNode, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list1
		}
		if list2 == nil {
			return list2
		}

		var curr *ListNode
		if list1.Val < list2.Val {
			curr = list1
			list1 = list1.Next
		} else {
			curr = list2
			list2 = list2.Next
		}

		dummy := &ListNode{}
		dummy.Next = curr

		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				curr.Next = list1
				list1 = list1.Next
			} else {
				curr.Next = list2
				list2 = list2.Next
			}
			curr = curr.Next
		}

		if list1 != nil {
			curr.Next = list1
		} else if list2 != nil {
			curr.Next = list2
		}

		return dummy.Next
	}

	dummy := &ListNode{}
	dummy.Next = lists[0]
	for i := 1; i < len(lists); i++ {
		p, q := dummy.Next, lists[i]
		if q == nil {
			continue
		}
		if p == nil {
			dummy.Next = q
			continue
		}
		dummy.Next = mergeTwoListsHelper(p, q)
	}

	return dummy.Next
}

// MIN HEAP IMPLEMENTATION

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int { return len(h) }

func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h ListNodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(node any) { *h = append(*h, node.(*ListNode)) }

func (h *ListNodeHeap) Pop() any {
	old := *h
	n := len(old)
	*h = (*h)[:n-1]
	return old[n-1]
}

func mergeKListsHeap(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := &ListNodeHeap{}
	heap.Init(h)
	for _, list := range lists {
		heap.Push(h, list)
	}

	dummy := &ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := heap.Pop(h)
		curr.Next = node.(*ListNode)
		curr = curr.Next
		if curr.Next != nil {
			heap.Push(h, curr.Next)
		}
	}
	return dummy.Next
}
