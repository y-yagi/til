package main

type ListNode struct {
	Val  int
	Next *ListNode
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
