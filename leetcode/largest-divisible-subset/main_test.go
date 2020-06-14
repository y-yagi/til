package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	want []int
}{
	{[]int{1, 2, 3}, []int{1, 2}},
	{[]int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
	{[]int{1, 2, 4, 8, 16}, []int{1, 2, 4, 8, 16}},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := largestDivisibleSubset(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in '%v', got '%v', want %v", tt.in1, got, tt.want)
		}
	}
}
