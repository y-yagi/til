package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var answer [][]int
	if root == nil {
		return answer
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)

	for len(q) > 0 {
		var tmp []int
		for _, tree := range q {
			q = q[1:]
			tmp = append(tmp, tree.Val)

			if tree.Left != nil {
				q = append(q, tree.Left)
			}
			if tree.Right != nil {
				q = append(q, tree.Right)
			}
		}

		answer = append(answer, tmp)
	}
	return answer
}
