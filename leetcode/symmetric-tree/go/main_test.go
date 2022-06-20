package main

import (
	"testing"
)

var tests = []struct {
	in   *TreeNode
	want bool
}{
	{buildTree([]*int{p(1), p(2), p(2), p(3), p(4), p(4), p(3)}), true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isSymmetric(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}

func buildTree(s []*int) *TreeNode {
	x, s := s[0], s[1:]
	root := &TreeNode{Val: *x}
	buildBranch(root, s)

	return root
}

func buildBranch(top *TreeNode, s []*int) {
	if len(s) == 0 {
		return
	}

	x, s := s[0], s[1:]
	left := &TreeNode{Val: *x}
	top.Left = left

	x, s = s[0], s[1:]
	right := &TreeNode{Val: *x}
	top.Right = right

	buildBranch(top.Left, s)
	buildBranch(top.Right, s)
}

func p(x int) *int {
	return &x
}
