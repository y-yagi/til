package main

import (
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 10, 5, 25, 2, 8}, 28},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := findMaximumXOR(tt.in1)
		if got != tt.want {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
