package main

import "testing"

var tests = []struct {
	in  string
	out bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
	{"]", false},
	{"(])", false},
	{"[(({})}]", false},
	{"(([]){})", true},
	{"{}{}()[]", true},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := isValid(tt.in)
		if got != tt.out {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.out)
		}
	}
}
