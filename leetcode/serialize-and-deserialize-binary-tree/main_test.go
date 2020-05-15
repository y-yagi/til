package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	root := &TreeNode{Val: 1}

	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Left = nil
	root.Left.Right = nil
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	obj := Constructor()

	data := obj.serialize(root)
	expected := "1,2,null,null,3,4,null,null,5,null,null,"
	if data != expected {
		t.Fatalf("got %v, want %v", data, expected)
	}

	ans := obj.deserialize(data)
	if root.Val != ans.Val {
		t.Fatalf("got %v, want %v", ans.Val, root.Val)
	}
}
