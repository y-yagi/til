package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p, q := headA, headB

	for p != nil || q != nil {
		if p == q {
			return p
		}
		if p == nil {
			p = headB
			q = q.Next
			continue
		}
		if q == nil {
			q = headA
			p = p.Next
			continue
		}
		p = p.Next
		q = q.Next
	}

	return nil
}
