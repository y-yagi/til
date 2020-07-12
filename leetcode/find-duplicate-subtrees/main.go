package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	cache := make(map[string]int)
	result := make([]*TreeNode, 0)
	find(root, cache, &result)
	return result
}

func find(root *TreeNode, cache map[string]int, result *[]*TreeNode) string {
	if root == nil {
		return "|"
	}

	left, right := find(root.Left, cache, result), find(root.Right, cache, result)
	key := strconv.Itoa(root.Val) + "|" + left + "|" + right
	if cache[key] == 1 {
		*result = append(*result, root)
	}
	cache[key]++
	return key
}
