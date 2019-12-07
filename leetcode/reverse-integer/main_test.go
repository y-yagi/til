package main

import "testing"

var tests = []struct {
	in  int
	out int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		got := reverse(tt.in)
		if got != tt.out {
			t.Fatalf("in %d, got %d, want %d", tt.in, got, tt.out)
		}
	}
}
