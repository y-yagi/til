package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return digDepth(root, 0)
}

func digDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	depth++
	leftDepth := depth
	rightDepth := depth

	if root.Left != nil {
		leftDepth = digDepth(root.Left, depth)
	}

	if root.Right != nil {
		rightDepth = digDepth(root.Right, depth)
	}

	if root.Left == nil {
		return rightDepth
	}

	if root.Right == nil {
		return leftDepth
	}

	if leftDepth > rightDepth {
		return rightDepth
	}

	return leftDepth
}
