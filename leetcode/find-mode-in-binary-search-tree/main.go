package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	count := make(map[int]int)
	answer := []int{}
	max := 0

	search(root, count)

	for _, v := range count {
		if v > max {
			max = v
		}
	}

	for k, v := range count {
		if v == max {
			answer = append(answer, k)
		}
	}

	return answer
}

func search(root *TreeNode, count map[int]int) {
	if root == nil {
		return
	}

	count[root.Val]++

	if root.Left != nil {
		search(root.Left, count)
	}
	if root.Right != nil {
		search(root.Right, count)
	}
}
