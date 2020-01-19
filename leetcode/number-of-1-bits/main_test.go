package main

import "testing"

var tests = []struct {
	in  uint32
	out int
}{
	{11, 3},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := hammingWeight(tt.in)
		if got != tt.out {
			t.Fatalf("in %d, got %d, want %d", tt.in, got, tt.out)
		}
	}
}
