package main

import "testing"

var tests = []struct {
	in1 string
	in2 string
	out int
}{
	{"mississippi", "issip", 4},
	{"aaaaa", "bba", -1},
	{"hello", "ll", 2},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, tt := range tests {
		got := strStr(tt.in1, tt.in2)
		if got != tt.out {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.out)
		}
	}
}
