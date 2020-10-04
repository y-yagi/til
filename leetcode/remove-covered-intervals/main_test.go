package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{1, 2}, []int{1, 4}, []int{3, 4}}, 1},
	{[][]int{[]int{1, 4}, []int{3, 6}, []int{2, 8}}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := removeCoveredIntervals(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
