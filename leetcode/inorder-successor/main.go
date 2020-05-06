package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	queue := make([]*TreeNode, 0)
	inorder := math.MinInt64

	for len(queue) != 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}

		root = queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]

		if inorder == p.Val {
			return root
		}

		inorder = root.Val
		root = root.Right
	}

	return nil
}
