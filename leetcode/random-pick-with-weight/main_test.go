package main

import (
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	s := Constructor([]int{1, 3})
	got := s.PickIndex()
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got '%v', want '%v'", got, want)
	}
}
