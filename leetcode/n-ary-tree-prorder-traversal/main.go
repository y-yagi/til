package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	stack := []*Node{root}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return ans
}
