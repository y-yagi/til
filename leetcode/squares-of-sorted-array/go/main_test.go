package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := sortedSquares(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
