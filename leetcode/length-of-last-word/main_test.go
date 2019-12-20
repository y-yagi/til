package main

import "testing"

var tests = []struct {
	in  string
	out int
}{
	{"b   a    ", 1},
	{"a ", 1},
	{"Hello world", 5},
	{"        ", 0},
	{"a", 1},
	{"Today is a nice day", 3},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, tt := range tests {
		got := lengthOfLastWord(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
