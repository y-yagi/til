package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 1, 3}, 1},
	{[]int{3, 3, 3, 1}, 1},
	{[]int{1, 1, 3, 1}, 1},
	{[]int{1, 1}, 1},
	{[]int{1, 3, 5}, 1},
	{[]int{2, 2, 2, 0, 1}, 0},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findMin(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
