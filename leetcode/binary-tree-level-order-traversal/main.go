package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		curr := []int{}
		for i := 0; i < l; i++ {
			n := queue[0]
			queue = queue[1:]

			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}

			curr = append(curr, n.Val)
		}

		ans = append(ans, curr)
	}

	return ans
}
