package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := []int{}
	stack2 := []int{}
	merged := []int{}

	curr := l1
	for curr != nil {
		stack1 = append([]int{curr.Val}, stack1...)
		curr = curr.Next
	}

	curr = l2
	for curr != nil {
		stack2 = append([]int{curr.Val}, stack2...)
		curr = curr.Next
	}

	l := max(len(stack1), len(stack2))
	carryon := 0
	for i := 0; i < l; i++ {
		t := carryon
		if i < len(stack1) {
			t += stack1[i]
		}
		if i < len(stack2) {
			t += stack2[i]
		}

		if t >= 10 {
			carryon = 1
			merged = append(merged, t-10)
		} else {
			merged = append(merged, t)
			carryon = 0
		}
	}

	if carryon != 0 {
		merged = append(merged, carryon)
	}

	dummy := &ListNode{}
	curr = dummy
	for i := len(merged) - 1; i >= 0; i-- {
		t := &ListNode{Val: merged[i]}
		curr.Next = t
		curr = t
	}

	return dummy.Next
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}
