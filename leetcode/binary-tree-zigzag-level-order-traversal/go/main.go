package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0

	for len(queue) > 0 {
		qSize := len(queue)
		tmp := make([]int, 0)

		for i := 0; i < qSize; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if level%2 == 1 {
			reverse(tmp)
		}
		result = append(result, tmp)
		level++
	}

	return result
}

func reverse(nums []int) {
	start, end := 0, len(nums)-1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
