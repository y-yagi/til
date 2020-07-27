package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  []int
	in2  []int
	want float64
}{
	{[]int{}, []int{1}, 1},
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestIsValid(t *testing.T) {
	for _, tt := range tests {
		got := findMedianSortedArrays(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
		}
	}
}
