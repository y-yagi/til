package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// With recursion.
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

// Without recursion(for avoid stack overflow).
func dfs2(root, target *Node) bool {
	visited := map[*Node]*Node{}
	stack := []*Node{root}

	for len(stack) != 0 {
		cur := stack[0]
		stack = stack[1:]

		if cur == target {
			return true
		}

		for _, next := range cur.Neighbors {
			if _, found := visited[next]; !found {
				visited[next] = next
				stack = append(stack, next)
			}
		}
	}

	return false
}
