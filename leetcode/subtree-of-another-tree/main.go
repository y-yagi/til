package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	nodes := []*TreeNode{s}
	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]
		if node == nil {
			continue
		}

		if same(node, t) {
			return true
		}
		nodes = append(nodes, node.Left)
		nodes = append(nodes, node.Right)
	}

	return false
}

func same(s, t *TreeNode) bool {
	leftResult := true
	rightResult := true
	if s.Val != t.Val {
		return false
	}

	if s.Left != nil {
		leftResult = false
		if t.Left == nil {
			return false
		}
		leftResult = same(s.Left, t.Left)
	} else if t.Left != nil {
		return false
	}

	if s.Right != nil {
		rightResult = false
		if t.Right == nil {
			return false
		}
		rightResult = same(s.Right, t.Right)
	} else if s.Right != nil {
		return false
	}

	return leftResult && rightResult
}
