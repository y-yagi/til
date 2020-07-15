package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	obj := Constructor([]string{"hello"})
	if !obj.IsUnique("hello") {
		t.Fatalf("fail!")
	}

	obj = Constructor([]string{"deer", "door", "cake", "card"})
	if obj.IsUnique("dear") {
		t.Fatalf("fail!")
	}
}
