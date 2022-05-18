package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{1, true},
	{16, true},
	{-16, false},
	{218, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isPowerOfTwo(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
