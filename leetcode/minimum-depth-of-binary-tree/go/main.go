package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	ans := math.MaxInt32
	if root.Left != nil {
		ans = min(minDepth(root.Left), ans)
	}
	if root.Right != nil {
		ans = min(minDepth(root.Right), ans)
	}

	return ans + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
