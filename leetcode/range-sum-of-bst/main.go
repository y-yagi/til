package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	dfs(root, low, high, &ans)
	return ans
}

func dfs(node *TreeNode, low int, high int, ans *int) {
	if node != nil {
		if low <= node.Val && node.Val <= high {
			*ans += node.Val
		}

		if node.Val > low {
			dfs(node.Left, low, high, ans)
		}

		if node.Val < high {
			dfs(node.Right, low, high, ans)
		}
	}
}
