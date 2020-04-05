package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  int
	want int
}{
	{[]int{1, 2, 3, 4}, 6, 1},
	{[]int{1, 2, 3, 4, 5, 6, 7, 1, 23, 21, 3, 1, 2, 1, 1, 1, 1, 1, 12, 2, 3, 2, 3, 2, 2}, 6, 5},
	{[]int{1, 2, 3}, 3, 2},
	{[]int{-1, -1, 1}, 0, 1},
	{[]int{1, 1, 1}, 2, 2},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := subarraySum(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
