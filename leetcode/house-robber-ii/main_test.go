package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{2, 3, 2}, 3},
	// {[]int{1, 2, 1, 1}, 3},
	// {[]int{1, 2, 3, 1}, 4},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := rob(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
