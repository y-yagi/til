package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, V int) []*TreeNode {
	if root == nil {
		return []*TreeNode{nil, nil}
	} else if root.Val <= V {
		bns := splitBST(root.Right, V)
		root.Right = bns[0]
		bns[0] = root
		return bns
	} else {
		bns := splitBST(root.Left, V)
		root.Left = bns[1]
		bns[1] = root
		return bns
	}
}
