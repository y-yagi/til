package main

import "testing"

var tests = []struct {
	in   []int
	want bool
}{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	{[]int{2, 1}, false},
	{[]int{3, 5, 5}, false},
	{[]int{0, 3, 2, 1}, true},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := validMountainArray(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
