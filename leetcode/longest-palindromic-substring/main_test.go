package main

import "testing"

var tests = []struct {
	in  string
	out string
}{
	{"babad", "aba"},
	{"caba", "aba"},
	{"cbbd", "bb"},
	{"a", "a"},
	{"ac", "c"},
	{"bb", "bb"},
	{"abb", "bb"},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		got := longestPalindrome(tt.in)
		if got != tt.out {
			t.Fatalf("in %q, got %q, want %q", tt.in, got, tt.out)
		}
	}
}
