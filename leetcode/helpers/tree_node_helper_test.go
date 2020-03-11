package main

import (
	"testing"
)

var tests = []struct {
	in   []*int
	want *TreeNode
}{
	{[]*int{Int(4), Int(2)}, nil},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := buildTreeNode(tt.in)
		if got == nil || got.Val != 4 {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
