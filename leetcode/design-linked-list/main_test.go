package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor()
	obj.Get(1)
	obj.AddAtHead(1)
	obj.AddAtTail(5)
	obj.AddAtIndex(2, 3)
	obj.DeleteAtIndex(1)
}
