package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	recuseTree(root, p, q)
	return root
}

func recuseTree(current, p, q *TreeNode) bool {
	if current == nil {
		return false
	}

	var left, right, mid int
	if recuseTree(current.Left, p, q) {
		left = 1
	}

	if recuseTree(current.Right, p, q) {
		right = 1
	}

	if current == p || current == q {
		mid = 1
	}

	if mid+left+right >= 2 {
		ans = current
	}

	return (mid+left+right > 0)
}
