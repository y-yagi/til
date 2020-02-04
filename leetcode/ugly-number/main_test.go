package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want bool
}{
	{6, true},
	{8, true},
	{14, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := isUgly(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
