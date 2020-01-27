package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(root *TreeNode) {
	if root == nil {
		return
	}

	n := root.Right
	root.Right = root.Left
	root.Left = n

	invert(root.Right)
	invert(root.Left)
}
