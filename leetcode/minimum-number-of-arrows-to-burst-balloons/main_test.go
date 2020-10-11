package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	want int
}{
	{[][]int{[]int{10, 16}, []int{2, 8}, []int{1, 6}, []int{7, 12}}, 2},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := findMinArrowShots(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
