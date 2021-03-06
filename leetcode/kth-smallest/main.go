package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	for true {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}
