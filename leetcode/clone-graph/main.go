package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, found := visited[node]; found {
		return visited[node]
	}
	cloneNode := &Node{Val: node.Val}
	visited[node] = cloneNode

	for _, neighbor := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, clone(neighbor, visited))
	}
	return cloneNode
}
