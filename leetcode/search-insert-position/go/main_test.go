package main

import "testing"

var tests = []struct {
	in1 []int
	in2 int
	out int
}{
	{[]int{1, 3}, 2, 1},
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, tt := range tests {
		got := searchInsert(tt.in1, tt.in2)
		if got != tt.out {
			t.Fatalf("in1 %v, in2 %v, got %v, want %v", tt.in1, tt.in2, got, tt.out)
		}
	}
}
