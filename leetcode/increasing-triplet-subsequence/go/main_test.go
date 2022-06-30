package main

import "testing"

var tests = []struct {
	in1  []int
	want bool
}{
	{[]int{1, 0, 10, 0, 0, 1000}, true},
	{[]int{1, 0, 0, 1}, false},
	{[]int{5, 1, 5, 5, 2, 5, 4}, true},
	{[]int{1, 2, 3, 4, 5}, true},
	{[]int{5, 4, 3, 2, 1}, false},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		got := increasingTriplet(tt.in1)
		if got != tt.want {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
