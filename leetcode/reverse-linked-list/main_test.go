package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	in1  *ListNode
	want *ListNode
}{
	{buildList([]int{1, 2, 3, 4, 5}), buildList([]int{5, 4, 3, 2, 1})},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := reverseList(tt.in1)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}
	}
}

func (x *ListNode) Equal(y *ListNode) bool {
	if x == nil || y == nil {
		return false
	}
	return x.Val == y.Val
}

func buildList(s []int) *ListNode {
	var result *ListNode
	var before *ListNode

	for _, v := range s {
		result = &ListNode{Val: v}
		if before != nil {
			result.Next = before
		}
		before = result
	}

	return result
}
