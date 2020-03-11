package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Int(v int) *int { return &v }

func buildTreeNode(values []*int) *TreeNode {
	var v *int
	v, values = values[0], values[1:]
	root := TreeNode{Val: *v}

	for len(values) > 0 {
		v, values = values[0], values[1:]

		if v == nil {
		}
	}

	return &root
}
