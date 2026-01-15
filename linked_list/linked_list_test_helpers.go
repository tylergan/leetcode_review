package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromSlice(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	curr := head
	for _, v := range values[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

func sliceFromList(head *ListNode) []int {
	var out []int
	for curr := head; curr != nil; curr = curr.Next {
		out = append(out, curr.Val)
	}
	return out
}
