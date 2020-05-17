package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor(3)
	obj.Move(0, 0, 1)
	obj.Move(0, 1, 1)
	got := obj.Move(0, 2, 1)
	if got != 1 {
		t.Fatalf("got '%v'!", got)
	}
}
