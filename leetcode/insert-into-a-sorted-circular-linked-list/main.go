package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		newNode := &Node{Val: x}
		newNode.Next = newNode
		return newNode
	}

	prev := aNode
	curr := aNode.Next
	insert := false

	for {
		if prev.Val <= x && x <= curr.Val {
			insert = true
		} else if prev.Val > curr.Val {
			if x >= prev.Val || x <= curr.Val {
				insert = true
			}
		}

		if insert {
			prev.Next = &Node{Val: x, Next: curr}
			return aNode
		}

		prev = curr
		curr = curr.Next

		if prev == aNode {
			break
		}
	}

	prev.Next = &Node{Val: x, Next: curr}
	return aNode
}
