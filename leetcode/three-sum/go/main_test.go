package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want [][]int
}{
	{[]int{3, 0, -2, -1, 1, 2}, [][]int{[]int{-2, -1, 3}, []int{-2, 0, 2}, []int{-1, 0, 1}}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := threeSum(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
