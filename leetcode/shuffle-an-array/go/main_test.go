package main

import (
	"reflect"
	"testing"
)

func TestIsValid1(t *testing.T) {
	nums := []int{1, 2, 3}
	obj := Constructor(nums)
	obj.Shuffle()
	got := obj.Reset()

	if !reflect.DeepEqual(nums, got) {
		t.Fatalf("got %v, want %v", got, nums)
	}
}
