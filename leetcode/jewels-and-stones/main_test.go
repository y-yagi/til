package main

import (
	"testing"
)

var tests = []struct {
	in1  string
	in2  string
	want int
}{
	{"aA", "aAAbbbb", 3},
	{"z", "ZZ", 0},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := numJewelsInStones(tt.in1, tt.in2)
		if got != tt.want {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
