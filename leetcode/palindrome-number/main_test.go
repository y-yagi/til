package main

import "testing"

var tests = []struct {
	in  int
	out bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{0, true},
	{1000021, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		got := isPalindrome(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
