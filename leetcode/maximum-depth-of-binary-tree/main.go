package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
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

	if leftDepth >= rightDepth {
		return leftDepth
	}
	return rightDepth

}
