package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	current := prev.Next

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			t := current.Val
			for current != nil && current.Val == t {
				current = current.Next
			}
			prev.Next = current
		} else {
			prev = current
			current = prev.Next
		}
	}

	return dummyHead.Next
}
