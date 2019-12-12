package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var tests = []struct {
	in1  *ListNode
	in2  *ListNode
	want *ListNode
}{
	{buildList([]int{1, 2, 4}), buildList([]int{1, 3, 4}), buildList([]int{1, 1, 2, 3, 4, 4})},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := mergeTwoLists(tt.in1, tt.in2)
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
