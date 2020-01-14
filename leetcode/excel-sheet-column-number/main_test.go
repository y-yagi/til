package main

import (
	"testing"
)

var tests = []struct {
	in   string
	want int
}{
	{"A", 1},
	{"AB", 28},
	{"AZ", 52},
	{"ZY", 701},
	{"AAA", 703},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := titleToNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
