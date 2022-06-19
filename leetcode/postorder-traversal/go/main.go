package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	dig(root, &ans)
	return ans
}

func dig(node *TreeNode, ans *[]int) {
	if node.Left != nil {
		dig(node.Left, ans)
	}

	if node.Right != nil {
		dig(node.Right, ans)
	}
	*ans = append(*ans, node.Val)

}
