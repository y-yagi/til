package main

import (
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	in3  int
	want bool
}{
	{[]int{1, 2, 3, 1}, 3, 0, true},
	{[]int{1, 0, 1, 1}, 1, 2, true},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := containsNearbyAlmostDuplicate(tt.in1, tt.in2, tt.in3)
		if got != tt.want {
			t.Fatalf("in %v - %v, got %v, want %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
