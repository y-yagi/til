package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validate(root, nil, nil)
}

func validate(node, row, high *TreeNode) bool {
	if node == nil {
		return true
	}

	if row != nil && node.Val <= row.Val {
		return false
	}

	if high != nil && node.Val >= high.Val {
		return false
	}

	return validate(node.Left, row, node) && validate(node.Right, node, high)
}
