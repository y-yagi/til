package main

import (
	"reflect"
	"testing"
)

func TestSuccess(t *testing.T) {
	obj := Constructor()
	got1 := obj.Ls("/")
	if !reflect.DeepEqual(got1, []string{}) {
		t.Fatalf("got '%v', want '%v'", got1, []string{})
	}

	obj.Mkdir("/test")
	obj.AddContentToFile("/test/a", "dummy")
	got2 := obj.ReadContentFromFile("/test/a")
	if !reflect.DeepEqual(got2, "dummy") {
		t.Fatalf("got '%v', want '%v'", got1, "dummy")
	}
}
