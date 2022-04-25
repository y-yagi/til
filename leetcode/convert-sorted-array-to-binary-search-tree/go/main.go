package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(0, len(nums)-1, nums)
}

func helper(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}

	p := (left + right) / 2

	// preorder traversal: node -> left -> right
	root := &TreeNode{Val: nums[p]}
	root.Left = helper(left, p-1, nums)
	root.Right = helper(p+1, right, nums)
	return root
}
