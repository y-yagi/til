package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		qSize := len(queue)

		for i := 0; i < qSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i < qSize-1 {
				node.Next = queue[0]
			} else {
				node.Next = nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
