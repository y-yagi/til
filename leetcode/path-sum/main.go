package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	sum -= root.Val

	if root.Right == nil && root.Left == nil {
		if sum == 0 {
			return true
		} else {
			return false
		}
	}

	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
