package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{3, 2, 3}, 3},
	{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := majorityElement(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
