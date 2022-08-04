package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	a := Constructor([]int{-2, 0, 3, -5, 2, -1})

	s := a.SumRange(0, 2)
	if s != 1 {
		t.Fatalf("got %v, want %v", s, 1)
	}

	s = a.SumRange(0, 5)
	if s != -3 {
		t.Fatalf("got %v, want %v", s, -3)
	}
}
