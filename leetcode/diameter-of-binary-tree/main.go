package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	helper(root, &ans)
	return ans - 1
}

func helper(node *TreeNode, ans *int) int {
	if node == nil {
		return 0
	}
	L := helper(node.Left, ans)
	R := helper(node.Right, ans)

	*ans = max(*ans, L+R+1)
	return max(L, R) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
