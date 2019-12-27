package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q != nil {
			return false
		}
		return true
	} else if q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	if p.Right != nil {
		if q.Right == nil {
			return false
		} else {
			if p.Right.Val != q.Right.Val {
				return false
			} else {
				answer := isSameTree(p.Right, q.Right)
				if answer == false {
					return false
				}
			}
		}
	} else if q.Right != nil {
		return false
	}

	if p.Left != nil {
		if q.Left == nil {
			return false
		} else {
			if p.Left.Val != q.Left.Val {
				return false
			} else {
				answer := isSameTree(p.Left, q.Left)
				if answer == false {
					return false
				}
			}
		}
	} else if q.Left != nil {
		return false
	}

	return true
}
