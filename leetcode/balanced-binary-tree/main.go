package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	v1 := float64(height(root.Left))
	v2 := float64(height(root.Right))

	if int(math.Abs(v1-v2)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	v1 := float64(height(node.Left))
	v2 := float64(height(node.Right))

	return 1 + int(math.Max(v1, v2))
}
