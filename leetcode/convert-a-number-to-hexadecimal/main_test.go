package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  int
	want string
}{
	{26, "1a"},
	{-1, "ffffffff"},
	{-2, "fffffffe"},
	{0, "0"},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := toHex(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
