package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	swap(dummyHead)
	return dummyHead.Next
}

func swap(node *ListNode) {
	if node.Next == nil || node.Next.Next == nil {
		return
	}

	x := node.Next
	y := node.Next.Next.Next
	node.Next = node.Next.Next
	node.Next.Next = x
	node.Next.Next.Next = y
	swap(node.Next.Next)
}
