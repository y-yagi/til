package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{2, false},
	{19, true},
	{100, true},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isHappy(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
