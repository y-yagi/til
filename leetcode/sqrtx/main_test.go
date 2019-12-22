package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want int
}{
	{4, 2},
	{8, 2},
}

func TestAddBinary(t *testing.T) {
	for _, tt := range tests {
		got := mySqrt(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
