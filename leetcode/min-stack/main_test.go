package main

import (
	"testing"
)

func TestIsValid1(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	x := obj.GetMin()
	if x != -3 {
		t.Fatalf("got %v, want %v", x, -3)
	}

	obj.Pop()

	x = obj.Top()
	if x != 0 {
		t.Fatalf("got %v, want %v", x, 0)
	}

	x = obj.GetMin()
	if x != -2 {
		t.Fatalf("got %v, want %v", x, -2)
	}
}
