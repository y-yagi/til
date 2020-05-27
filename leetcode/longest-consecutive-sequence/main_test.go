package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	{[]int{1, 2, 0, 1}, 3},
	{[]int{0, -1}, 2},
	{[]int{0}, 1},
	{[]int{100, 4, 200, 1, 3, 2}, 4},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := longestConsecutive(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
