package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(3)
	obj.Add(5)
	if result := obj.Find(4); !result {
		t.Fatalf("fail!")
	}

	obj = Constructor()
	obj.Add(0)
	obj.Add(0)
	obj.Add(0)
	if result := obj.Find(0); !result {
		t.Fatalf("fail!")
	}
}
