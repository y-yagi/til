package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string

	if root == nil {
		return paths
	}

	searchPath(root, "", &paths)
	return paths
}

func searchPath(root *TreeNode, path string, paths *[]string) {
	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
		return
	}

	if root.Left != nil {
		searchPath(root.Left, path+"->", paths)
	}

	if root.Right != nil {
		searchPath(root.Right, path+"->", paths)
	}
}
