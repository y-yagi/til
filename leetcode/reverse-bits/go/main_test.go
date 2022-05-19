package main

import "testing"

var tests = []struct {
	in  uint32
	out uint32
}{
	{43261596, 964176192},
}

func TestReverseBits(t *testing.T) {
	for _, tt := range tests {
		got := reverseBits(tt.in)
		if got != tt.out {
			t.Fatalf("in %d, got %d, want %d", tt.in, got, tt.out)
		}
	}
}
