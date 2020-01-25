package main

import (
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want bool
}{
	{"egg", "add", true},
	{"foo", "bar", false},
	{"paper", "title", true},
	{"ab", "aa", false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isIsomorphic(tt.in1, tt.in2)
		if got != tt.want {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
