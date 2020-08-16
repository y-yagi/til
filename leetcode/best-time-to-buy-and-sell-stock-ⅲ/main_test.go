package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 13},
	{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
	{[]int{1, 2, 3, 4, 5}, 4},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestMaxProfit(t *testing.T) {
	for _, tt := range tests {
		got := maxProfit(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
