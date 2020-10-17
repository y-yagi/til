package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1  [][]int
	in2  int
	want bool
}{
	{[][]int{[]int{1}, []int{3}}, 3, true},
	{[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}, 3, true},
	{[][]int{[]int{1, 3, 5, 7}, []int{10, 11, 16, 20}, []int{23, 30, 34, 50}}, 13, false},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
		got := searchMatrix(tt.in1, tt.in2)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 '%v', got '%v', want '%v'", tt.in1, got, tt.want)
		}
	}
}
