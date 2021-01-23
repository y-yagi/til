package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := l1
	q := l2
	curr := dummyHead
	carry := 0

	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
		curr = curr.Next
	}

	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
