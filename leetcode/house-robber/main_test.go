package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{1, 2, 3, 1}, 4},
	{[]int{2, 7, 9, 3, 1}, 12},
	{[]int{1, 2}, 2},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := rob(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
