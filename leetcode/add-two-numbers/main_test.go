package main

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	l1_3 := ListNode{Val: 3}
	l1_2 := ListNode{Val: 4, Next: &l1_3}
	l1 := ListNode{Val: 2, Next: &l1_2}

	l2_3 := ListNode{Val: 4}
	l2_2 := ListNode{Val: 6, Next: &l2_3}
	l2 := ListNode{Val: 5, Next: &l2_2}

	got := addTwoNumbers(&l1, &l2)
	if !reflect.DeepEqual(got.Val, 7) {
		t.Fatalf("got %v, want %v", got.Val, 7)
	}

	if !reflect.DeepEqual(got.Next.Val, 0) {
		t.Fatalf("got %v, want %v", got.Next.Val, 0)
	}

	if !reflect.DeepEqual(got.Next.Next.Val, 8) {
		t.Fatalf("got %v, want %v", got.Next.Next.Val, 8)
	}

	l1_3 = ListNode{Val: 9}
	l1_2 = ListNode{Val: 4, Next: &l1_3}
	l1 = ListNode{Val: 2, Next: &l1_2}

	l2_4 := ListNode{Val: 9}
	l2_3 = ListNode{Val: 4, Next: &l2_4}
	l2_2 = ListNode{Val: 6, Next: &l2_3}
	l2 = ListNode{Val: 5, Next: &l2_2}

	got = addTwoNumbers(&l1, &l2)
	if !reflect.DeepEqual(got.Val, 7) {
		t.Fatalf("got %v, want %v", got.Val, 7)
	}

	if !reflect.DeepEqual(got.Next.Val, 0) {
		t.Fatalf("got %v, want %v", got.Next.Val, 0)
	}

	if !reflect.DeepEqual(got.Next.Next.Val, 4) {
		t.Fatalf("got %v, want %v", got.Next.Next.Val, 4)
	}
}
