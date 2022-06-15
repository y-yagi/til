package main

import (
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want bool
}{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
	{"ab", "a", false},
	{"aacc", "ccac", false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isAnagram(tt.in1, tt.in2)
		if got != tt.want {
			t.Fatalf("in1 '%v', in2 '%v', got '%v', want '%v'", tt.in1, tt.in2, got, tt.want)
		}
	}
}
