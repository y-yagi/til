package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want int
}{
	{10, 4},
	{3, 1},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := countPrimes(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
