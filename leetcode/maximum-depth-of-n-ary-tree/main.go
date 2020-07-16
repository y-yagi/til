package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var answer [][]int

	q := []*Node{root}

	for len(q) > 0 {
		var tmp []int
		for _, tree := range q {
			q = q[1:]
			tmp = append(tmp, tree.Val)

			for _, node := range tree.Children {
				q = append(q, node)
			}
		}

		answer = append(answer, tmp)
	}

	return len(answer)
}
