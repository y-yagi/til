package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var prev *ListNode

	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}

	con := prev
	tail := cur
	var third *ListNode
	for n > 0 {
		third = cur.Next
		cur.Next = prev
		prev = cur
		cur = third
		n--
	}

	if con != nil {
		con.Next = prev
	} else {
		head = prev
	}

	tail.Next = cur
	return head
}
