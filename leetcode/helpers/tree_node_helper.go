package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Int(v int) *int { return &v }

func buildTree(values []*int) *TreeNode {
	var value *int

	root := &TreeNode{}
	node := &TreeNode{}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(values) > 0 {
		qSize := len(queue)

		for i := 0; i < qSize; i++ {
			if len(values) == 0 {
				return root
			}

			queue, node = queue[1:], queue[0]
			values, value = values[1:], values[0]

			if value == nil {
				node = nil
			} else {
				node.Val = *value

				if len(queue) < len(values) {
					left := &TreeNode{}
					node.Left = left
					queue = append(queue, left)
				}

				if len(queue) < len(values) {
					right := &TreeNode{}
					node.Right = right
					queue = append(queue, right)
				}
			}
		}
	}

	return root
}
