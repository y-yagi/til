package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return builder(0, 0, len(inorder)-1, preorder, inorder)
}

func builder(preStart, inStart, inEnd int, preorder, inorder []int) *TreeNode {
	if preStart > len(preorder)-1 || inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	inIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			inIndex = i
		}
	}
	root.Left = builder(preStart+1, inStart, inIndex-1, preorder, inorder)
	root.Right = builder(preStart+inIndex-inStart+1, inIndex+1, inEnd, preorder, inorder)
	return root
}
