package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{3, 0, 1}, 2},
	{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	{[]int{1}, 0},
	{[]int{0}, 1},
	{[]int{0, 1}, 2},
	{[]int{1, 0}, 2},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := missingNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
