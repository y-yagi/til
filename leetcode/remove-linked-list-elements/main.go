package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	tl := &ListNode{}

	head, tl.Next = tl, head

	for tl.Next != nil {
		if tl.Next.Val == val {
			tl.Next = tl.Next.Next
		} else {
			tl = tl.Next
		}
	}
	return head.Next
}
