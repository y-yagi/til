package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return dig(root, 0)
}

func dig(node *TreeNode, depth int) int {
	if node == nil {
		return depth
	}

	return max(dig(node.Left, depth+1), dig(node.Right, depth+1))
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
