package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{root.Val}
	queue := []*TreeNode{root.Right, root.Left}
	for len(queue) != 0 {
		l := len(queue)
		added := false
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				queue = append(queue, node.Right)
				queue = append(queue, node.Left)
			}

			if node != nil && !added {
				ans = append(ans, node.Val)
				added = true
			}
		}
	}

	return ans
}
