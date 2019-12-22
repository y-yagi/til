package main

import (
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want string
}{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
}

func TestAddBinary(t *testing.T) {
	for _, tt := range tests {
		got := addBinary(tt.in1, tt.in2)
		if got != tt.want {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
