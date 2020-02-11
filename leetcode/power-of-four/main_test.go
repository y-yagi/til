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
	{0, false},
	{5, false},
	{-2147483648, false},
	{2, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isPowerOfFour(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
