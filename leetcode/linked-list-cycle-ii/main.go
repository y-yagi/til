package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := map[*ListNode]bool{}

	for head != nil {
		if _, found := visited[head]; found {
			return head
		}

		visited[head] = true
		head = head.Next
	}

	return nil
}
