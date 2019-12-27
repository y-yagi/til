package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	in1  *TreeNode
	in2  *TreeNode
	want bool
}{
	{buildTree([]int{1, 2, 3}), buildTree([]int{1, 2, 3}), true},
	{buildTree([]int{1, 2, 1}), buildTree([]int{1, 1, 2}), false},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isSameTree(tt.in1, tt.in2)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}
	}
}

func (x *TreeNode) Equal(y *TreeNode) bool {
	if x == nil || y == nil {
		return false
	}

	if x.Val != y.Val {
		return false
	}

	if x.Right != nil {
		if y.Right == nil {
			return false
		}

		if x.Right.Val != y.Right.Val {
			return false
		}
	}

	if x.Left != nil {
		if y.Left == nil {
			return false
		}

		if x.Left.Val != y.Left.Val {
			return false
		}
	}

	return true
}

func buildTree(s []int) *TreeNode {
	var top *TreeNode

	top = &TreeNode{Val: s[0]}
	if len(s) > 1 {
		left := &TreeNode{Val: s[1]}
		top.Left = left
	}
	if len(s) > 2 {
		right := &TreeNode{Val: s[2]}
		top.Right = right
	}
	return top
}
