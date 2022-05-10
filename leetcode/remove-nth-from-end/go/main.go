package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	first := dummy
	second := dummy

	for i := 1; i <= n+1; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
