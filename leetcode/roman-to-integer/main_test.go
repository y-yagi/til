package main

import "testing"

var tests = []struct {
	in  string
	out int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, tt := range tests {
		got := romanToInt(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
