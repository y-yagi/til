package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	pseudoHead := &ListNode{}
	curr := head
	var prevNode, nextNode *ListNode

	for curr != nil {
		prevNode = pseudoHead
		nextNode = pseudoHead.Next

		for nextNode != nil {
			if curr.Val < nextNode.Val {
				break
			}
			prevNode = nextNode
			nextNode = nextNode.Next
		}

		nextIter := curr.Next
		curr.Next = nextNode
		prevNode.Next = curr

		curr = nextIter
	}

	return pseudoHead.Next
}
