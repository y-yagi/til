package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	newNode := &Node{Val: x}
	if aNode == nil {
		newNode.Next = newNode
		return newNode
	}

	curr := aNode.Next
	prev := aNode

	for curr != nil && curr != aNode {
		if prev.Val <= x && curr.Val >= x {
			prev.Next = newNode
			newNode.Next = curr
			return aNode
		} else if prev.Val > curr.Val {
			if x >= prev.Val || x <= curr.Val {
				prev.Next = newNode
				newNode.Next = curr
				return aNode
			}
		}

		prev = curr
		curr = curr.Next
	}

	prev.Next = newNode
	newNode.Next = curr

	return aNode
}
