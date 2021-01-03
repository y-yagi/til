package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	data := [][]int{}
	dfs(root, 0, 0, &data)
	sort.Slice(data, func(i, j int) bool {
		if data[i][0] == data[j][0] {
			if data[i][1] == data[j][1] {
				return data[i][2] < data[j][2]
			}
			return data[i][1] > data[j][1]
		}
		return data[i][0] < data[j][0]
	})
	lastX := data[0][0]
	traversal := [][]int{{}}
	for _, v := range data {
		if v[0] != lastX {
			traversal = append(traversal, []int{})
			lastX = v[0]
		}
		traversal[len(traversal)-1] = append(traversal[len(traversal)-1], v[2])
	}
	return traversal
}

func dfs(node *TreeNode, x, y int, data *[][]int) {
	if node == nil {
		return
	}
	*data = append(*data, []int{x, y, node.Val})
	dfs(node.Left, x-1, y-1, data)
	dfs(node.Right, x+1, y-1, data)
}
