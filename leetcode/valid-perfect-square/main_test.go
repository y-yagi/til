package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{16, true},
	{14, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isPerfectSquare(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got, '%v', want '%v'", tt.in, got, tt.want)
		}
	}
}
