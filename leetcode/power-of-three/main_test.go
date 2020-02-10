package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{27, true},
	{0, false},
	{9, true},
	{45, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isPowerOfThree(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
