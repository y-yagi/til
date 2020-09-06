package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ans := []int{}
	queue := []*TreeNode{root1, root2}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		ans = append(ans, node.Val)
	}

	sort.Ints(ans)
	return ans
}
