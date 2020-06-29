package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generate(1, n)
}
func generate(start, end int) []*TreeNode {
	allTrees := []*TreeNode{}
	if start > end {
		allTrees = append(allTrees, nil)
		return allTrees
	}

	for i := start; i <= end; i++ {
		leftTrees := generate(start, i-1)
		rightTrees := generate(i+1, end)

		for _, lNode := range leftTrees {
			for _, rNode := range rightTrees {
				current := &TreeNode{Val: i}
				current.Left = lNode
				current.Right = rNode
				allTrees = append(allTrees, current)
			}
		}
	}
	return allTrees
}
