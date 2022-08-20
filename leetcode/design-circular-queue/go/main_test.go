package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor(3)
	obj.EnQueue(1)
	obj.EnQueue(2)
	obj.EnQueue(3)
	obj.EnQueue(4)
	obj.Rear()
	obj.IsFull()
	obj.DeQueue()
	obj.EnQueue(4)
	obj.Rear()
}
