package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}

	stack := [][]*TreeNode{}
	stack = append(stack, []*TreeNode{t1, t2})

	for len(stack) > 0 {
		t := stack[0]
		stack = stack[1:]

		if t[0] == nil || t[1] == nil {
			continue
		}

		t[0].Val += t[1].Val
		if t[0].Left == nil {
			t[0].Left = t[1].Left
		} else {
			stack = append(stack, []*TreeNode{t[0].Left, t[1].Left})
		}

		if t[0].Right == nil {
			t[0].Right = t[1].Right
		} else {
			stack = append(stack, []*TreeNode{t[0].Right, t[1].Right})
		}
	}

	return t1
}
