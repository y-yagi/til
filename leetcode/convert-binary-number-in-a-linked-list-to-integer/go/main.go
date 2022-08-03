package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	num := head.Val
	for head.Next != nil {
		num = num*2 + head.Next.Val
		// num = (num << 1) | head.next.val;
		head = head.Next
	}

	return num
}
