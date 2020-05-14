package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	iterator := Constructor([][]int{[]int{1, 2}, []int{3}, []int{4}})
	if got := iterator.Next(); got != 1 {
		t.Fatalf("got %v, want %v", got, 1)
	}

	if got := iterator.Next(); got != 2 {
		t.Fatalf("got %v, want %v", got, 2)
	}

	if got := iterator.Next(); got != 3 {
		t.Fatalf("got %v, want %v", got, 3)
	}

	if got := iterator.HasNext(); got != true {
		t.Fatalf("got %v, want %v", got, true)
	}

	if got := iterator.Next(); got != 4 {
		t.Fatalf("got %v, want %v", got, 4)
	}

	if got := iterator.HasNext(); got != false {
		t.Fatalf("got %v, want %v", got, false)
	}
}

func TestEmpty(t *testing.T) {
	iterator := Constructor([][]int{})
	if got := iterator.HasNext(); got != false {
		t.Fatalf("got %v, want %v", got, false)
	}
}

func TestOne(t *testing.T) {
	iterator := Constructor([][]int{[]int{1}, []int{}})
	if got := iterator.HasNext(); got != true {
		t.Fatalf("got %v, want %v", got, true)
	}

	if got := iterator.Next(); got != 1 {
		t.Fatalf("got %v, want %v", got, 1)
	}

	if got := iterator.HasNext(); got != false {
		t.Fatalf("got %v, want %v", got, false)
	}
}
