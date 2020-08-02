package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{0}, 0},
	{[]int{0, 9}, 9},
	{[]int{2, 1, 5, 6, 2, 3}, 10},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := largestRectangleArea(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
