package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	got := obj.Get(1)
	want := 1
	if got != want {
		t.Fatalf("got '%v', want '%v'", got, want)
	}

	obj.Put(3, 3)
	got = obj.Get(2)
	want = -1
	if got != want {
		t.Fatalf("got '%v', want '%v'", got, want)
	}
}
