package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	calcurate(root, &ans)
	return ans
}

func calcurate(node *TreeNode, sum *int) int {
	if node == nil {
		return 0
	}

	leftGain := max(calcurate(node.Left, sum), 0)
	rightGain := max(calcurate(node.Right, sum), 0)
	priceNewpath := node.Val + leftGain + rightGain
	*sum = max(*sum, priceNewpath)

	return node.Val + max(leftGain, rightGain)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
