package main

import "testing"

var tests = []struct {
	in  string
	out int
}{
	{"abccccdd", 7},
	{"ccc", 3},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		got := longestPalindrome(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
