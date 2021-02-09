package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		t := curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}

	return prev
}
