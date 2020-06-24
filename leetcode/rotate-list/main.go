package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oldTail := head
	n := 0
	for n = 1; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	newTail := head
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
