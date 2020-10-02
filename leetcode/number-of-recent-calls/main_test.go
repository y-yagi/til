package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor()
	obj.Ping(1)
	obj.Ping(100)
	obj.Ping(3001)
	got := obj.Ping(3002)
	want := 3
	if got != want {
		t.Fatalf("got '%v', want '%v'", got, want)
	}
}
