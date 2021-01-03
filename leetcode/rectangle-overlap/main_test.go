package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	want bool
}{
	{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
	{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
	{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := isRectangleOverlap(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
