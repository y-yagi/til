package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if leftToRight {
				level = append(level, node.Val)
			} else {
				level = append([]int{node.Val}, level...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, level)
		leftToRight = !leftToRight
	}

	return ans
}
