package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	in   *ListNode
	want *ListNode
}{
	{buildList([]int{1, 1, 2}), buildList([]int{1, 2})},
	{buildList([]int{1, 1, 2, 2, 3, 3}), buildList([]int{1, 2, 3})},
	{buildList([]int{1, 1, 2, 3, 3}), buildList([]int{1, 2, 3})},
	{buildList([]int{1, 1, 1}), buildList([]int{1})},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, tt := range tests {
		got := deleteDuplicates(tt.in)
		if diff := cmp.Diff(tt.want, got); diff != "" {
			t.Fatalf("mismatch (-want +got):\n%s", diff)
		}
	}
}

func (x *ListNode) Equal(y *ListNode) bool {
	for {
		if x == nil && y == nil {
			return true
		}

		if x == nil || y == nil {
			return false
		}

		if x.Val != y.Val {
			return false
		}

		x = x.Next
		y = y.Next
	}

	return true
}

func buildList(s []int) *ListNode {
	var result *ListNode
	var before *ListNode

	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	for _, v := range s {
		result = &ListNode{Val: v}
		if before != nil {
			result.Next = before
		}
		before = result
	}

	return result
}
