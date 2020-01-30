package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var prev *ListNode = nil
	var slow *ListNode = head
	var fast *ListNode = head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	if fast != nil {
		slow = slow.Next
	}
	var head2 *ListNode = slow
	head = reverse(head)

	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head

	for curr != nil {
		var post *ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = post
	}

	return prev
}
