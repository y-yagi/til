package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
func main() {
	l1_3 := ListNode{Val: 3}
	l1_2 := ListNode{Val: 4, Next: &l1_3}
	l1 := ListNode{Val: 2, Next: &l1_2}

	l2_3 := ListNode{Val: 4}
	l2_2 := ListNode{Val: 6, Next: &l2_3}
	l2 := ListNode{Val: 5, Next: &l2_2}

	l := addTwoNumbers(&l1, &l2)
	fmt.Printf("%v %v %v\n", l.Val, l.Next.Val, l.Next.Next.Val)

	l3_4 := ListNode{Val: 1}
	l3_3 := ListNode{Val: 0, Next: &l3_4}
	l3_2 := ListNode{Val: 0, Next: &l3_3}
	l3_1 := ListNode{Val: 0, Next: &l3_2}
	l3 := ListNode{Val: 1, Next: &l3_1}

	l = addTwoNumbers(&l3, &l2)
	fmt.Printf("%v %v %v %v %v\n", l.Val, l.Next.Val, l.Next.Next.Val, l.Next.Next.Next.Val, l.Next.Next.Next.Next.Val)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1s, l2s string
	var result *ListNode

	l := l1
	for {
		l1s = strconv.Itoa(l.Val) + l1s
		if l.Next == nil {
			break
		}
		l = l.Next
	}

	l = l2
	for {
		l2s = strconv.Itoa(l.Val) + l2s
		if l.Next == nil {
			break
		}
		l = l.Next
	}

	l1i, _ := new(big.Int).SetString(l1s, 0)
	l2i, _ := new(big.Int).SetString(l2s, 0)
	sum := l1i.Add(l1i, l2i)

	s := strings.Split(fmt.Sprintf("%v", sum), "")
	rs := ""
	for _, v := range s {
		rs = string(v) + rs
	}

	for i := len(rs) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(string(rs[i]))
		now := &ListNode{Val: val}

		if i == len(rs)-1 {
			result = now
		} else {
			now.Next = result
			result = now
		}
	}

	return result
}
