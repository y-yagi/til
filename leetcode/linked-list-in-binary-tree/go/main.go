package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return findSubPath(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func findSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	return head.Val == root.Val && (findSubPath(head.Next, root.Left) || findSubPath(head.Next, root.Right))
}
