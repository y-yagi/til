package main

func countComponents(n int, edges [][]int) int {
	roots := make([]int, n)
	for i := 0; i < n; i++ {
		roots[i] = i
	}

	for _, e := range edges {
		root1 := find(roots, e[0])
		root2 := find(roots, e[1])
		if root1 != root2 {
			roots[root1] = root2
			n--
		}
	}
	return n
}

func find(roots []int, id int) int {
	for roots[id] != id {
		roots[id] = roots[roots[id]]
		id = roots[id]
	}
	return id
}
