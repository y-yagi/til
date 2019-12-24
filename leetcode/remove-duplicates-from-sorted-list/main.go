package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	now := head.Next
	before := head

	for {
		if now.Val == before.Val {
			before.Next = now.Next
		} else {
			before = now
		}

		if now.Next == nil {
			break
		}

		now = now.Next
	}

	return head
}
