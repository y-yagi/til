package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	var ans int
	minValue := root.Val
	maxValue := root.Val
	dig(root, maxValue, minValue, &ans)
	return ans
}

func dig(node *TreeNode, maxValue, minValue int, ans *int) {
	*ans = max(*ans, maxValue-node.Val)
	*ans = max(*ans, node.Val-minValue)
	maxValue = max(maxValue, node.Val)
	minValue = min(minValue, node.Val)

	if node.Left != nil {
		dig(node.Left, maxValue, minValue, ans)
	}

	if node.Right != nil {
		dig(node.Right, maxValue, minValue, ans)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
