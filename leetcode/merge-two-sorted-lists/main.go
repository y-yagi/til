package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	s := []int{}

	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			s = append(s, l1.Val)
			l1 = l1.Next
		}

		if l2 != nil {
			s = append(s, l2.Val)
			l2 = l2.Next
		}
	}

	return buildList(s)
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
