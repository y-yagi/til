package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want int
}{
	{3, 0},
	{5, 1},
	{7, 1},
	{30, 7},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := trailingZeroes(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
