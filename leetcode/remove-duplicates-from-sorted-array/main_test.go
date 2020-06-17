package main

import "testing"

var tests = []struct {
	in   []int
	want int
}{
	{[]int{1, 1}, 1},
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range tests {
		got := removeDuplicates(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
