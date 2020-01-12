package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want string
}{
	{1, "A"},
	{26, "Z"},
	{28, "AB"},
	{52, "AZ"},
	{701, "ZY"},
	{703, "AAA"},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := convertToTitle(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
