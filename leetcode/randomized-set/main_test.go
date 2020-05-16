package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor()

	obj.Insert(1)
	obj.Remove(2)
	obj.Insert(2)
	obj.GetRandom()
	obj.Remove(1)
	obj.Insert(2)
	got := obj.GetRandom()
	if got != 2 {
		t.Fatalf("got '%v'!", got)
	}
}
