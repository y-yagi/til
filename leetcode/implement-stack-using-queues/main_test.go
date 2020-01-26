package main

import (
	"testing"
)

func TestIsValid1(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)

	x := stack.Top()
	if x != 2 {
		t.Fatalf("got %v, want %v", x, 2)
	}

	x = stack.Pop()
	if x != 2 {
		t.Fatalf("got %v, want %v", x, 2)
	}
}
