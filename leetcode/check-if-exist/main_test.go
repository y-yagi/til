package main

import "testing"

var tests = []struct {
	in   []int
	want bool
}{
	{[]int{4, -7, 11, 4, 18}, false},
	{[]int{-20, 8, -6, -14, 0, -19, 14, 4}, true},
	{[]int{10, 2, 5, 3}, true},
	{[]int{7, 1, 14, 11}, true},
	{[]int{3, 1, 7, 11}, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := checkIfExist(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
