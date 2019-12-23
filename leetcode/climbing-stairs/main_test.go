package main

import (
	"testing"
)

var tests = []struct {
	in   int
	want int
}{
	{2, 2},
	{3, 3},
	{4, 5},
}

func TestClimbStairs(t *testing.T) {
	for _, tt := range tests {
		got := climbStairs(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
