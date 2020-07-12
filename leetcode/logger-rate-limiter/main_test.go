package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	obj := Constructor()
	if result := obj.ShouldPrintMessage(1, "test"); !result {
		t.Fatalf("fail!")
	}

	if result := obj.ShouldPrintMessage(1, "test"); result {
		t.Fatalf("fail!!")
	}

	if result := obj.ShouldPrintMessage(11, "test"); !result {
		t.Fatalf("fail!!!")
	}
}
