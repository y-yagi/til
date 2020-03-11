package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var min int
	var vals []int

	vals = append(vals, root.Val)
	digDiff(root, &min, vals)
	return min
}

func digDiff(root *TreeNode, diff *int, vals []int) {
	if root == nil {
		return
	}

	checkMin(diff, root.Val, vals)
	vals = append(vals, root.Val)

	if root.Left != nil {
		digDiff(root.Left, diff, vals)
	}

	if root.Right != nil {
		digDiff(root.Right, diff, vals)
	}
}

func checkMin(diff *int, val int, vals []int) {
	for _, x := range vals {
		abs := int(math.Abs(float64(val - x)))
		if *diff == 0 {
			*diff = abs
		} else if *diff > abs {
			*diff = abs
		}
	}
}
