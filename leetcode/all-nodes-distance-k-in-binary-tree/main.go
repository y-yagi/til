package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var parents map[*TreeNode]*TreeNode

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parents = make(map[*TreeNode]*TreeNode)
	dfs(root, nil)

	queue := []*TreeNode{target}
	seen := make(map[*TreeNode]bool)
	seen[target] = true

	dist := 0

	for len(queue) > 0 {

		size := len(queue)
		if dist == K {

			ans := []int{}
			for _, node := range queue {
				ans = append(ans, node.Val)
			}

			return ans
		}

		for i := 0; i < size; i++ {

			popped := queue[0]
			queue = queue[1:]

			if popped.Left != nil && !seen[popped.Left] {
				seen[popped.Left] = true
				queue = append(queue, popped.Left)
			}

			if popped.Right != nil && !seen[popped.Right] {
				seen[popped.Right] = true
				queue = append(queue, popped.Right)
			}

			p := parents[popped]

			if p != nil && !seen[p] {
				seen[p] = true
				queue = append(queue, p)
			}
		}

		dist++
	}

	return nil
}

func dfs(node, parent *TreeNode) {
	if node != nil {
		parents[node] = parent
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
}
