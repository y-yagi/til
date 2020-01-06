package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
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
