package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	return sums(root, 0, sum, map[int]int{0: 1})
}

func sums(node *TreeNode, curSum, target int, m map[int]int) int {
	if node == nil {
		return 0
	}

	curSum += node.Val
	summary := m[curSum-target]
	m[curSum]++
	summary += sums(node.Left, curSum, target, m)
	summary += sums(node.Right, curSum, target, m)
	m[curSum]--
	return summary
}
