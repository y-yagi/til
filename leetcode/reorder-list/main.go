package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// find the middle of linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second part of the list
	var prev *ListNode
	curr := slow
	for curr != nil {
		t := curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}

	// merge two sorted linked lists
	first := head
	second := prev
	for second.Next != nil {
		t := first.Next
		first.Next = second
		first = t

		t = second.Next
		second.Next = first
		second = t
	}
}
