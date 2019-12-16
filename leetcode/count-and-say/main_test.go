package main

import "testing"

var tests = []struct {
	in  int
	out string
}{
	{1, "1"},
	{4, "1211"},
	{6, "312211"},
}

func TestCountAndSay(t *testing.T) {
	for _, tt := range tests {
		got := countAndSay(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
