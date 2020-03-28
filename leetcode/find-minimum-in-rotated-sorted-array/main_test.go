package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want int
}{
	{[]int{3, 4, 5, 1, 2}, 1},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findMin(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
