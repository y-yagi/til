package main

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor(3)
	obj.Next(3)
	obj.Next(10)
}
