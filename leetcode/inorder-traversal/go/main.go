package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return dig(root, []int{})
}

func dig(node *TreeNode, values []int) []int {
	if node == nil {
		return values
	}

	if node.Left != nil {
		values = dig(node.Left, values)
	}
	values = append(values, node.Val)
	if node.Right != nil {
		values = dig(node.Right, values)
	}

	return values
}
