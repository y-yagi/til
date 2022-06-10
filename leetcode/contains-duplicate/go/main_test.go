package main

import (
	"testing"
)

var tests = []struct {
	in   []int
	want bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
}

func TestValid(t *testing.T) {
	for _, tt := range tests {
		got := containsDuplicate(tt.in)
		if got != tt.want {
			t.Fatalf("in %v, got %v, want %v", tt.in, got, tt.want)
		}
	}
}
