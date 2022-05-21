package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 2, 1}, 0},
	{[]int{1, 2}, 1},
	{[]int{1}, 0},
	{[]int{2, 1}, 0},
	{[]int{1, 2, 3, 1}, 2},
	{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findPeakElement(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
