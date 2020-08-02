package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}

	var first, last *Node
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node != nil {
			dfs(node.Left)

			if last != nil {
				last.Right = node
				node.Left = last
			} else {
				first = node
			}

			last = node

			dfs(node.Right)
		}
	}

	dfs(root)

	last.Right = first
	first.Left = last

	return first
}
