package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor()
	for i := 1; i <= 3; i++ {
		obj.AddNum(i)
	}

	if obj.FindMedian() != 2.0 {
		t.Fatalf("got '%v', want '%v'", obj.FindMedian(), 2.0)
	}
}
