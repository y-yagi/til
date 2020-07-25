package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest, val := root.Val, root.Val
	for root != nil {
		val = root.Val
		if math.Abs(float64(val)-target) < math.Abs(float64(closest)-target) {
			closest = val
		} else {
			closest = closest
		}

		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closest
}
