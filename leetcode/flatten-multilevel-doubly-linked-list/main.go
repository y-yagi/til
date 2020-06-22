package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	pseudoHead := &Node{0, nil, root, nil}
	var curr, prev *Node = nil, pseudoHead

	stack := []*Node{}
	stack = append(stack, root)

	for len(stack) > 0 {
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prev.Next = curr
		curr.Prev = prev

		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}
		prev = curr
	}

	pseudoHead.Next.Prev = nil
	return pseudoHead.Next
}
