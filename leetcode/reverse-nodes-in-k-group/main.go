package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var next = head
	for i := 0; i < k; i++ {
		if next != nil {
			next = next.Next
		} else {
			return head
		}
	}

	var originHead = head
	var newHead *ListNode = nil
	var nextElem = head.Next
	for head.Next != next {
		nextElem, head.Next = head.Next, newHead
		newHead, head = head, nextElem
	}
	head.Next = newHead
	originHead.Next = reverseKGroup(next, k)

	return head
}
