package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	cur := head
	ahead := cur
	for ahead != nil && ahead.Next != nil {
		cur = cur.Next
		ahead = ahead.Next.Next
	}

	return cur
}
