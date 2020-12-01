package main

import (
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor("leetcode.com")
	obj.Visit("google.com")
	obj.Visit("facebook.com")
	obj.Visit("youtube.com")
	obj.Back(1)
	obj.Back(1)
	obj.Forward(1)
	obj.Visit("linkedin.com")
	obj.Forward(2)

	got := obj.Back(2)
	want := "google.com"
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got '%v', want '%v'", got, want)
	}

	obj.Back(7)
}
