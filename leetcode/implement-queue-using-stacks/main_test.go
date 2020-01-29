package main

import (
	"testing"
)

func TestIsValid1(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)

	x := queue.Peek()
	if x != 1 {
		t.Fatalf("got %v, want %v", x, 1)
	}

	x = queue.Pop()
	if x != 1 {
		t.Fatalf("got %v, want %v", x, 1)
	}
}
