package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func dfs(cur, target *Node, visited map[*Node]*Node) bool {
	if cur == target {
		return true
	}

	for _, next := range cur.Neighbors {
		if _, found := visited[next]; !found {
			visited[next] = next
			if dfs(next, target, visited) {
				return true
			}
		}
	}

	return false
}
