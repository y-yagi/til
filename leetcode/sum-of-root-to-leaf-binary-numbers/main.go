package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	rootToLeaf := 0
	preordre(root, 0, &rootToLeaf)
	return rootToLeaf
}

func preordre(node *TreeNode, currNumber int, rootToLeaf *int) {
	if node != nil {
		currNumber = (currNumber << 1) | node.Val
		if node.Left == nil && node.Right == nil {
			*rootToLeaf += currNumber
		}
		preordre(node.Left, currNumber, rootToLeaf)
		preordre(node.Right, currNumber, rootToLeaf)
	}
}
