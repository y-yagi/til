package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	prev := newHead
	curr := prev.Next
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			t := curr.Val
			for curr != nil && curr.Val == t {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = curr
			curr = prev.Next
		}
	}

	return newHead.Next
}
