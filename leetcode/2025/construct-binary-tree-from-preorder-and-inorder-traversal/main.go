package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Find the index of the root in inorder array
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// Build left and right subtrees
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}
